// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Four string

const (
	FourStop          Four = "stop"
	FourLength        Four = "length"
	FourContentFilter Four = "content_filter"
)

func (e Four) ToPointer() *Four {
	return &e
}
func (e *Four) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stop":
		fallthrough
	case "length":
		fallthrough
	case "content_filter":
		*e = Four(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Four: %v", v)
	}
}

type MessageChunk struct {
	ChatSessionID *string `json:"chat_session_id,omitempty"`
	ContentChunk  string  `json:"content_chunk"`
	FinishReason  any     `json:"finish_reason,omitempty"`
}

func (o *MessageChunk) GetChatSessionID() *string {
	if o == nil {
		return nil
	}
	return o.ChatSessionID
}

func (o *MessageChunk) GetContentChunk() string {
	if o == nil {
		return ""
	}
	return o.ContentChunk
}

func (o *MessageChunk) GetFinishReason() any {
	if o == nil {
		return nil
	}
	return o.FinishReason
}
