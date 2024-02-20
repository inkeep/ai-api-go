// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/inkeep/ai-api-go/internal/utils"
)

type AssistantMessage struct {
	role         string        `const:"assistant" json:"role"`
	Content      string        `json:"content"`
	RecordsCited *RecordsCited `json:"records_cited,omitempty"`
}

func (a AssistantMessage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssistantMessage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AssistantMessage) GetRole() string {
	return "assistant"
}

func (o *AssistantMessage) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *AssistantMessage) GetRecordsCited() *RecordsCited {
	if o == nil {
		return nil
	}
	return o.RecordsCited
}
