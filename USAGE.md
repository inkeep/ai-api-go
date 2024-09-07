<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	aiapigo "github.com/inkeep/ai-api-go"
	"github.com/inkeep/ai-api-go/models/components"
	"log"
)

func main() {
	s := aiapigo.New(
		aiapigo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.ChatSession.Create(ctx, components.CreateChatSessionWithChatResultInput{
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
	})
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