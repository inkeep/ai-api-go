# ChatSession
(*ChatSession*)

## Overview

Create and manage chat sessions for users. Chat history and continuation of chat is automatically done.

### Available Operations

* [Create](#create) - Create Chat Session
* [Continue](#continue) - Continue Chat Session

## Create

Create Chat Session

### Example Usage

```go
package main

import(
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"context"
	"log"
)

func main() {
    s := aiapigo.New(
        aiapigo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    request := components.CreateChatSessionWithChatResultInput{
        IntegrationID: "<value>",
        ChatSession: components.ChatSessionInput{
            Messages: []components.Message{
                components.CreateMessageUserMessage(
                    components.UserMessage{
                        Content: "<value>",
                    },
                ),
            },
        },
    }
    ctx := context.Background()
    res, err := s.ChatSession.Create(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.ChatResult != nil {
        defer res.ChatResultStream.Close()

        for res.ChatResultStream.Next() {
            event := res.ChatResultStream.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [components.CreateChatSessionWithChatResultInput](../../models/components/createchatsessionwithchatresultinput.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |


### Response

**[*operations.CreateResponse](../../models/operations/createresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.HTTPValidationError | 422                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

## Continue

Continue Chat Session

### Example Usage

```go
package main

import(
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"context"
	"log"
)

func main() {
    s := aiapigo.New(
        aiapigo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )
    var chatSessionID string = "<value>"

    continueChatSessionWithChatResultInput := components.ContinueChatSessionWithChatResultInput{
        IntegrationID: "<value>",
        Message: components.CreateMessageAssistantMessage(
                components.AssistantMessage{
                    Content: "<value>",
                },
        ),
    }
    ctx := context.Background()
    res, err := s.ChatSession.Continue(ctx, chatSessionID, continueChatSessionWithChatResultInput)
    if err != nil {
        log.Fatal(err)
    }
    if res.ChatResult != nil {
        defer res.ChatResultStream.Close()

        for res.ChatResultStream.Next() {
            event := res.ChatResultStream.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `chatSessionID`                                                                                                        | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `continueChatSessionWithChatResultInput`                                                                               | [components.ContinueChatSessionWithChatResultInput](../../models/components/continuechatsessionwithchatresultinput.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |


### Response

**[*operations.ContinueResponse](../../models/operations/continueresponse.md), error**
| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.HTTPValidationError | 422                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |
