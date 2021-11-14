package rollbar

import (
	go_rollbar "github.com/rollbar/rollbar-go"
)

// Config is a struct for configuration of Rollbar.
type Config struct {
	Token       string `yaml:"token" env:"ROLLBAR_TOKEN"`
	Environment string `yaml:"environment" env:"ROLLBAR_ENVIRONMENT"`
}

// Client is a struct for storing Rollbar client.
type Client struct {
	*go_rollbar.Client
}

// NewClient returns a new Rollbar client
func NewClient(config *Config, revision string) *Client {
	codeVersion := revision
	serverHost := ""
	serverRoot := ""
	client := go_rollbar.New(config.Token, config.Environment, codeVersion, serverHost, serverRoot)
	// ローカルやテスト環境などTokenがない場合はRollbarを無効にしておく
	client.SetEnabled(config.Token != "")
	return &Client{client}
}
