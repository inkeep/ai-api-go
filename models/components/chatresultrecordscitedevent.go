// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/inkeep/ai-api-go/internal/utils"
)

// ChatResultRecordsCitedEvent - A server-sent event with information about the records cited in the message.
type ChatResultRecordsCitedEvent struct {
	event string       `const:"records_cited" json:"event"`
	Data  RecordsCited `json:"data"`
}

func (c ChatResultRecordsCitedEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ChatResultRecordsCitedEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ChatResultRecordsCitedEvent) GetEvent() string {
	return "records_cited"
}

func (o *ChatResultRecordsCitedEvent) GetData() RecordsCited {
	if o == nil {
		return RecordsCited{}
	}
	return o.Data
}
