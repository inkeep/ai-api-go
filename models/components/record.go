// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package components

type Record struct {
	// The type of record
	Type        RecordType `json:"type"`
	URL         *string    `json:"url,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Breadcrumbs []string   `json:"breadcrumbs,omitempty"`
}

func (o *Record) GetType() RecordType {
	if o == nil {
		return RecordType{}
	}
	return o.Type
}

func (o *Record) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Record) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Record) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Record) GetBreadcrumbs() []string {
	if o == nil {
		return nil
	}
	return o.Breadcrumbs
}
