// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package aiapigo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/inkeep/ai-api-go/internal/hooks"
	"github.com/inkeep/ai-api-go/internal/utils"
	"github.com/inkeep/ai-api-go/models/components"
	"github.com/inkeep/ai-api-go/models/operations"
	"github.com/inkeep/ai-api-go/models/sdkerrors"
	"github.com/inkeep/ai-api-go/types/stream"
	"io"
	"net/http"
	"net/url"
)

// ChatSession - Create and manage chat sessions for users. Chat history and continuation of chat is automatically done.
type ChatSession struct {
	sdkConfiguration sdkConfiguration
}

func newChatSession(sdkConfig sdkConfiguration) *ChatSession {
	return &ChatSession{
		sdkConfiguration: sdkConfig,
	}
}

// Create Chat Session
func (s *ChatSession) Create(ctx context.Context, request components.CreateChatSessionWithChatResultInput, opts ...operations.Option) (*operations.CreateResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "create",
		SecuritySource: s.sdkConfiguration.Security,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/v0/chat_sessions/chat_results")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, text/event-stream;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"422", "4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out components.ChatResult
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out, ""); err != nil {
				return nil, err
			}

			res.ChatResult = &out
		case utils.MatchContentType(contentType, `text/event-stream`):
			out := stream.NewEventStream(httpRes.Body, func(se []byte) (components.ChatResultStream, error) {
				var e components.ChatResultStream
				if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(se), &e, ""); err != nil {
					return components.ChatResultStream{}, err
				}
				return e, nil
			})
			res.ChatResultStream = out
		default:
			rawBody, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.HTTPValidationError
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			rawBody, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := io.ReadAll(httpRes.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}

		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

// Continue Chat Session
func (s *ChatSession) Continue(ctx context.Context, chatSessionID string, continueChatSessionWithChatResultInput components.ContinueChatSessionWithChatResultInput, opts ...operations.Option) (*operations.ContinueResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "continue",
		SecuritySource: s.sdkConfiguration.Security,
	}

	request := operations.ContinueRequest{
		ChatSessionID:                          chatSessionID,
		ContinueChatSessionWithChatResultInput: continueChatSessionWithChatResultInput,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/v0/chat_sessions/{chat_session_id}/chat_results", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "ContinueChatSessionWithChatResultInput", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, text/event-stream;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"422", "4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ContinueResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out components.ChatResult
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out, ""); err != nil {
				return nil, err
			}

			res.ChatResult = &out
		case utils.MatchContentType(contentType, `text/event-stream`):
			out := stream.NewEventStream(httpRes.Body, func(se []byte) (components.ChatResultStream, error) {
				var e components.ChatResultStream
				if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(se), &e, ""); err != nil {
					return components.ChatResultStream{}, err
				}
				return e, nil
			})
			res.ChatResultStream = out
		default:
			rawBody, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out sdkerrors.HTTPValidationError
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out, ""); err != nil {
				return nil, err
			}
			return nil, &out
		default:
			rawBody, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		rawBody, err := io.ReadAll(httpRes.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}

		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
