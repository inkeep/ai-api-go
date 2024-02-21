// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/inkeep/ai-api-go/internal/utils"
)

type ContinueChatSessionWithChatResultInput struct {
	IntegrationID string  `json:"integration_id"`
	Message       Message `json:"message"`
	Stream        *bool   `default:"false" json:"stream"`
}

func (c ContinueChatSessionWithChatResultInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ContinueChatSessionWithChatResultInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ContinueChatSessionWithChatResultInput) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *ContinueChatSessionWithChatResultInput) GetMessage() Message {
	if o == nil {
		return Message{}
	}
	return o.Message
}

func (o *ContinueChatSessionWithChatResultInput) GetMessageUser() *UserMessage {
	return o.GetMessage().UserMessage
}

func (o *ContinueChatSessionWithChatResultInput) GetMessageAssistant() *AssistantMessage {
	return o.GetMessage().AssistantMessage
}

func (o *ContinueChatSessionWithChatResultInput) GetStream() *bool {
	if o == nil {
		return nil
	}
	return o.Stream
}