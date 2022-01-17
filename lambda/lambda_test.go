package lambda

import (
	"reflect"
	"testing"
)

var (
	testUserDetails = []*UserDetail{
		{
			UserId: "123",
			Email:  "something@email.com",
		},
		{
			UserId: "456",
			Email:  "nothing@email.com",
		},
		{
			UserId: "789",
			Email:  "everything@email.com",
		},
		{
			UserId: "012",
			Email:  "anything@email.com",
		},
	}
	goodLink = "https://example.com/this/shouldnt/work/at/all/24"
	badLink  = "https://example.com/this/shouldnt/work/at/all/24"
)

// func TestHandleSlackRequest(t *testing.T) {
// 	type args struct {
// 		ctx     context.Context
// 		request events.APIGatewayProxyRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := HandleSlackRequest(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
// 				t.Errorf("HandleSlackRequest() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func Test_pickRandomUsers(t *testing.T) {
	type args struct {
		userDetails   []*UserDetail
		reviewerCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test that reviewerCount limits the output.",
			args: args{
				userDetails:   testUserDetails,
				reviewerCount: 2,
			},
			want: 2,
		},
		{
			name: "Test that reviewerCount greater than the length of the users just returns all available users.",
			args: args{
				userDetails:   testUserDetails,
				reviewerCount: 10,
			},
			want: 4,
		},
		{
			name: "Test that reviewerCount of 0 returns an empty list.",
			args: args{
				userDetails:   testUserDetails,
				reviewerCount: 10,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickRandomUsers(tt.args.userDetails, tt.args.reviewerCount); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("pickRandomUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_getGroupDetails(t *testing.T) {
// 	type args struct {
// 		botCtx     slacker.BotContext
// 		group      string
// 		filterUser string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []*UserDetail
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := getGroupDetails(tt.args.botCtx, tt.args.group, tt.args.filterUser); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("getGroupDetails() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_updatePullRequest(t *testing.T) {
	type args struct {
		link        string
		userDetails []*UserDetail
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executePullRequestUpdate(tt.args.link, tt.args.userDetails)
			if (err != nil) != tt.wantErr {
				t.Errorf("updatePullRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updatePullRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBitbucketSlug(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test that getBitbucketSlug returns the correct string for a good link.",
			args: args{
				link: goodLink,
			},
			want: "example tbd; this fails right now",
		},
		{
			name: "Test that getBitbucketSlug returns a null string for a bad link",
			args: args{
				link: badLink,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBitbucketSlug(tt.args.link); got != tt.want {
				t.Errorf("getBitbucketSlug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBitbucketId(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test that getBitbucketId returns the correct string id for a good link.",
			args: args{
				link: goodLink,
			},
			want: "24",
		},
		{
			name: "Test that getBitbucketId returns a null string id for a bad link",
			args: args{
				link: badLink,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBitbucketId(tt.args.link); got != tt.want {
				t.Errorf("getBitbucketId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUserIds(t *testing.T) {
	type args struct {
		userDetails []*UserDetail
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test that getUserIds returns the correct list of string ids.",
			args: args{
				userDetails: testUserDetails,
			},
			want: []string{"123", "456", "789", "012"},
		},
		{
			name: "Test that getUserIds returns an empty list if passed an empty list",
			args: args{
				userDetails: []*UserDetail{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserIds(tt.args.userDetails); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUserIds() = %v, want %v", got, tt.want)
			}
		})
	}
}
