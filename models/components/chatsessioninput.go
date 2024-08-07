// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ChatSessionInput struct {
	Guidance *string   `json:"guidance,omitempty"`
	Context  *string   `json:"context,omitempty"`
	Messages []Message `json:"messages"`
	Tags     []string  `json:"tags,omitempty"`
}

func (o *ChatSessionInput) GetGuidance() *string {
	if o == nil {
		return nil
	}
	return o.Guidance
}

func (o *ChatSessionInput) GetContext() *string {
	if o == nil {
		return nil
	}
	return o.Context
}

func (o *ChatSessionInput) GetMessages() []Message {
	if o == nil {
		return []Message{}
	}
	return o.Messages
}

func (o *ChatSessionInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}
