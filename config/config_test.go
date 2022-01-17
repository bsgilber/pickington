package config

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	debug   = "true"
	env     = "local"
	secret  = "f897afsd9d779d"
	channel = "random-reviewer-channel"
)

var envMap = map[string]string{
	"SLACK_SECRET":     secret,
	"SLACK_CHANNEL_ID": channel,
	"DEBUG":            debug,
	"ENVIRONMENT":      env,
}

func init() {
	for k, v := range envMap {
		err := os.Setenv(k, v)
		if err != nil {
			log.Fatalf("could not set environment variable %s=%s: %v", k, v, err)
		}
	}
}

func TestConfig(t *testing.T) {
	cfg := NewConfig()
	assert.NotNil(t, cfg)
	assert.Equal(t, cfg.SlackAuthToken, secret)
	assert.Equal(t, cfg.SlackChannelId, channel)
	assert.Equal(t, debug, cfg.Debug)
}
