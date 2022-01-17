package config

import (
	"os"
)

const (
	slackBotTokenEnv = "SLACK_BOT_TOKEN"
	slackAppTokenEnv = "SLACK_APP_TOKEN"
	slackSecretEnv   = "SLACK_SECRET"
	slackChannelEnv  = "SLACK_CHANNEL_ID"
	bitbucketUserEnv = "BITBUCKET_USER"
	bitbucketPassEnv = "BITBUCKET_PASSWORD"
	debugEnv         = "DEBUG"
	envEnv           = "ENVIRONMENT"
)

// Config for iterable-lambda.
type Config struct {
	SlackBotToken  string
	SlackAppToken  string
	SlackAuthToken string
	SlackChannelId string
	BitbucketUser  string
	BitbucketPass  string
	Debug          string
}

// NewConfig initializes a new Config instance.
func NewConfig() *Config {
	return &Config{
		SlackBotToken:  os.Getenv(slackBotTokenEnv),
		SlackAppToken:  os.Getenv(slackAppTokenEnv),
		SlackChannelId: os.Getenv(slackChannelEnv),
		SlackAuthToken: os.Getenv(slackSecretEnv),
		BitbucketUser:  os.Getenv(bitbucketUserEnv),
		BitbucketPass:  os.Getenv(bitbucketPassEnv),
		Debug:          os.Getenv(debugEnv),
	}
}
