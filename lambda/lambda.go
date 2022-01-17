package lambda

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	b64 "encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/ktrysmt/go-bitbucket"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

type UserDetail struct {
	UserId string
	Email  string
}

const (
	timeout = 1 * time.Second
	retries = 2
)

var (
	transport = http.DefaultTransport.(*http.Transport).Clone()
	client    *httpclient.Client
)

func init() {
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	client = httpclient.NewClient(httpclient.WithHTTPTimeout(timeout), httpclient.WithRetryCount(retries), httpclient.WithHTTPClient(&http.Client{
		Timeout:   timeout,
		Transport: transport,
	}))
}

// HandleSlackRequest handles a Lambda event triggered by a APIGateway
func HandleSlackRequest(ctx context.Context, request events.APIGatewayProxyRequest) error {
	body, err := b64.URLEncoding.DecodeString(request.Body)
	if err != nil {
		panic(err)
	}

	bot := slacker.NewClient(cfg.SlackBotToken, cfg.SlackAppToken)
	threadReplyDefinition := &slacker.CommandDefinition{
		Description: "Tests errors in threads",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			sourceUser := botCtx.Event().User

			parsedBody, err := url.ParseQuery(string(body))
			if err != nil {
				panic(err)
			}

			codeReviewLink := parsedBody.Get("codereview")
			group := parsedBody.Get("group")
			reviewerCount, err := strconv.Atoi(parsedBody.Get("count"))
			if err != nil {
				panic(err)
			}

			userDetails := getGroupDetails(botCtx, group, sourceUser)
			shortList := pickRandomUsers(userDetails, reviewerCount)

			_, prErr := executePullRequestUpdate(codeReviewLink, shortList)
			if prErr != nil {
				panic(err)
			}

			response.Reply("succes")
		},
	}

	bot.Command("thread", threadReplyDefinition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	botErr := bot.Listen(ctx)
	if botErr != nil {
		log.Fatal(botErr)
	}

	return nil
}

func pickRandomUsers(userDetails []*UserDetail, reviewerCount int) []*UserDetail {
	userCount := len(userDetails)

	var limit int
	if userCount > reviewerCount {
		limit = reviewerCount
	} else {
		limit = userCount
	}

	reviewerList := make([]*UserDetail, len(userDetails))

	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(userCount)

	for i, v := range perm {
		reviewerList[v] = userDetails[i]
	}

	return reviewerList[:limit]
}

func getGroupDetails(botCtx slacker.BotContext, group string, filterUser string) []*UserDetail {
	client := botCtx.Client()
	members, err := client.GetUserGroupMembers(group)
	if err != nil {
		panic(err)
	}

	var userDetails []*UserDetail
	for _, member := range members {
		if member != filterUser {
			user, err := client.GetUserInfo(member)
			if err != nil {
				panic(err)
			}

			profile, err := client.GetUserProfile(&slack.GetUserProfileParameters{user.ID, false})
			if err != nil {
				panic(err)
			}
			userDetails = append(userDetails, &UserDetail{user.ID, profile.Email})
		}
	}

	return userDetails
}

func executePullRequestUpdate(link string, userDetails []*UserDetail) (interface{}, error) {
	c := bitbucket.NewBasicAuth(cfg.BitbucketUser, cfg.BitbucketPass)

	opt := &bitbucket.PullRequestsOptions{
		ID:                getBitbucketId(link),
		RepoSlug:          getBitbucketSlug(link),
		CloseSourceBranch: true,
		Reviewers:         getUserIds(userDetails),
	}

	return c.Repositories.PullRequests.Update(opt)
}

func getBitbucketSlug(link string, company string) string {
	re := regexp.MustCompile(`(?<=\/company\/)(.+)(?=\/pull-requests)`)
	return re.FindString(link)
}

func getBitbucketId(link string) string {
	re := regexp.MustCompile(`(?<=pull-requests\/)(\d+)`)
	return re.FindString(link)
}

func getUserIds(userDetails []*UserDetail) []string {
	ids := make([]string, len(userDetails))

	for i := range userDetails {
		ids[i] = userDetails[i].UserId
	}

	return ids
}
