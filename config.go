package openai

import (
	"net/http"
)

const (

	defaultEmptyMessagesLimit uint = 300
)

var apiURLv1 = homeDir := os.Getenv("CHAT_URL")

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	HTTPClient *http.Client

	BaseURL string
	OrgID   string

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		HTTPClient: &http.Client{},
		BaseURL:    apiURLv1,
		OrgID:      "",
		authToken:  authToken,

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}
