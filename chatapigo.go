// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package aiapigo

import (
	"context"
	"fmt"
	"github.com/inkeep/ai-api-go/internal/hooks"
	"github.com/inkeep/ai-api-go/internal/utils"
	"github.com/inkeep/ai-api-go/models/components"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production server
	"https://api.inkeep.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// ChatAPIGo - Inkeep Search and Chat API: Leverage the Inkeep APIs to create your own AI-powered search and chat experiences built on top of your own content.
type ChatAPIGo struct {
	// Create and manage chat sessions for users. Chat history and continuation of chat is automatically done.
	ChatSession *ChatSession

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*ChatAPIGo)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *ChatAPIGo) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *ChatAPIGo) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *ChatAPIGo) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *ChatAPIGo) {
		sdk.sdkConfiguration.Client = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiKey string) SDKOption {
	return func(sdk *ChatAPIGo) {
		security := components.Security{APIKey: apiKey}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *ChatAPIGo) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *ChatAPIGo) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *ChatAPIGo {
	sdk := &ChatAPIGo{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.1.0",
			SDKVersion:        "0.3.6",
			GenVersion:        "2.314.0",
			UserAgent:         "speakeasy-sdk/go 0.3.6 2.314.0 0.1.0 github.com/inkeep/ai-api-go",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.ChatSession = newChatSession(sdk.sdkConfiguration)

	return sdk
}
