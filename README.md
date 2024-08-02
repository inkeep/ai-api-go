# github.com/inkeep/ai-api-go

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/inkeep/ai-api-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"log"
	"os"
)

func main() {
	s := aiapigo.New(
		aiapigo.WithSecurity(os.Getenv("API_KEY")),
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [ChatSession](docs/sdks/chatsession/README.md)

* [Create](docs/sdks/chatsession/README.md#create) - Create Chat Session
* [Continue](docs/sdks/chatsession/README.md#continue) - Continue Chat Session
<!-- End Available Resources and Operations [operations] -->

<!-- Start Server-sent event streaming [eventstream] -->
## Server-sent event streaming

[Server-sent events][mdn-sse] are used to stream content from certain
operations. These operations will expose the stream as an iterable that
can be consumed using a simple `for` loop. The loop will
terminate when the server no longer has any events to send and closes the
underlying connection.

```go
package main

import (
	"context"
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"log"
	"os"
)

func main() {
	s := aiapigo.New(
		aiapigo.WithSecurity(os.Getenv("API_KEY")),
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

[mdn-sse]: https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
<!-- End Server-sent event streaming [eventstream] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.HTTPValidationError | 422                           | application/json              |
| sdkerrors.SDKError            | 4xx-5xx                       | */*                           |

### Example

```go
package main

import (
	"context"
	"errors"
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"github.com/inkeep/ai-api-go/models/sdkerrors"
	"log"
	"os"
)

func main() {
	s := aiapigo.New(
		aiapigo.WithSecurity(os.Getenv("API_KEY")),
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

		var e *sdkerrors.HTTPValidationError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.inkeep.com` | None |

#### Example

```go
package main

import (
	"context"
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"log"
	"os"
)

func main() {
	s := aiapigo.New(
		aiapigo.WithServerIndex(0),
		aiapigo.WithSecurity(os.Getenv("API_KEY")),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"log"
	"os"
)

func main() {
	s := aiapigo.New(
		aiapigo.WithServerURL("https://api.inkeep.com"),
		aiapigo.WithSecurity(os.Getenv("API_KEY")),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name        | Type        | Scheme      |
| ----------- | ----------- | ----------- |
| `APIKey`    | http        | HTTP Bearer |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"log"
	"os"
)

func main() {
	s := aiapigo.New(
		aiapigo.WithSecurity(os.Getenv("API_KEY")),
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
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"github.com/inkeep/ai-api-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	s := aiapigo.New(
		aiapigo.WithSecurity(os.Getenv("API_KEY")),
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
	res, err := s.ChatSession.Create(ctx, request, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
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

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"github.com/inkeep/ai-api-go/retry"
	"log"
	"os"
)

func main() {
	s := aiapigo.New(
		aiapigo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		aiapigo.WithSecurity(os.Getenv("API_KEY")),
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
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
