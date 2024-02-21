// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/inkeep/ai-api-go/internal/utils"
)

type MessageType string

const (
	MessageTypeUser      MessageType = "user"
	MessageTypeAssistant MessageType = "assistant"
)

type Message struct {
	UserMessage      *UserMessage
	AssistantMessage *AssistantMessage

	Type MessageType
}

func CreateMessageUser(user UserMessage) Message {
	typ := MessageTypeUser

	return Message{
		UserMessage: &user,
		Type:        typ,
	}
}

func CreateMessageAssistant(assistant AssistantMessage) Message {
	typ := MessageTypeAssistant

	return Message{
		AssistantMessage: &assistant,
		Type:             typ,
	}
}

func (u *Message) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Role string `json:"role"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Role {
	case "user":
		userMessage := new(UserMessage)
		if err := utils.UnmarshalJSON(data, &userMessage, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.UserMessage = userMessage
		u.Type = MessageTypeUser
		return nil
	case "assistant":
		assistantMessage := new(AssistantMessage)
		if err := utils.UnmarshalJSON(data, &assistantMessage, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.AssistantMessage = assistantMessage
		u.Type = MessageTypeAssistant
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Message) MarshalJSON() ([]byte, error) {
	if u.UserMessage != nil {
		return utils.MarshalJSON(u.UserMessage, "", true)
	}

	if u.AssistantMessage != nil {
		return utils.MarshalJSON(u.AssistantMessage, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}