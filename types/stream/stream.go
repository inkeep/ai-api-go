// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package stream

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type ServerEvent struct {
	ID    *string         `json:"id,omitempty"`
	Event *string         `json:"event,omitempty"`
	Data  json.RawMessage `json:"data,omitempty"`
	Retry *int64          `json:"retry,omitempty"`
}

var (
	boundary   = regexp.MustCompile(`\r\n\r\n|\r\r|\n\n`)
	lineEnding = regexp.MustCompile(`\r?\n|\r`)
)

func scanServerEvents(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	result := boundary.FindIndex(data)
	if result != nil {
		return result[1], data[:result[0]], nil
	}

	if atEOF {
		return len(data), bytes.TrimRight(data, "\r\n"), nil
	}

	return 0, nil, nil
}

type EventType interface {
	GetEventEncoding(event string) (string, error)
}

type EventStream[T any] struct {
	r            io.ReadCloser
	scanner      *bufio.Scanner
	unmarshaller func(se []byte) (T, error)
	sentinel     string

	finished bool
	err      error
	val      *T
}

func NewEventStream[T any](
	source io.Reader,
	unmarshaller func(se []byte) (T, error),
	sentinel string,
) *EventStream[T] {
	scanner := bufio.NewScanner(source)
	scanner.Split(scanServerEvents)

	var src io.ReadCloser
	if s, ok := source.(io.ReadCloser); ok {
		src = s
	} else {
		src = io.NopCloser(source)
	}

	return &EventStream[T]{
		r:            src,
		scanner:      scanner,
		unmarshaller: unmarshaller,
		sentinel:     sentinel,
	}
}

// Next waits for the next event from a stream which will be available
// through the Value() method. It returns false when the stream is done or
// an error occurred. After this method returns false, the Err method is used
// to check for any errors that occurred while parsing the stream.
func (es *EventStream[T]) Next() bool {
	if es.err != nil || es.finished {
		return false
	}

	if !es.scanner.Scan() {
		return false
	}

	es.err = es.scanner.Err()
	if es.err != nil {
		return false
	}

	b := es.scanner.Bytes()

	var event ServerEvent
	lines := lineEnding.Split(string(b), -1)
	publish := false
	data := ""
	for _, line := range lines {
		if line == "" {
			continue
		}

		delim := strings.Index(line, ":")
		if delim == 0 {
			continue
		}

		field := ""
		value := ""
		if delim > 0 {
			field = line[:delim]
		}
		if delim > 0 && delim < len(line)-1 {
			value = line[delim+1:]
		}
		value = strings.TrimPrefix(value, " ")

		switch field {
		case "id":
			publish = true
			event.ID = &value
		case "event":
			publish = true
			event.Event = &value
		case "retry":
			retry, err := strconv.ParseInt(value, 10, 64)
			if err == nil {
				publish = true
				event.Retry = &retry
			}
		case "data":
			publish = true
			data += value + "\n"
		}
	}

	if es.sentinel != "" && data == es.sentinel+"\n" {
		es.finished = true
		return false
	}

	if len(data) > 0 {
		data = data[:len(data)-1]
	}

	encoding := "application/json"

	var t T
	if et, ok := any(t).(EventType); ok {
		ev := ""
		if event.Event != nil {
			ev = *event.Event
		}

		var err error
		encoding, err = et.GetEventEncoding(ev)
		if err != nil {
			es.err = err
			return false
		}
	} else {
		var a interface{}
		if err := json.Unmarshal([]byte(data), &a); err != nil {
			encoding = "string"
		}
	}

	if encoding == "string" {
		data = fmt.Sprintf("%q", data)
	}

	event.Data = []byte(data)

	e, err := json.Marshal(event)
	if err != nil {
		es.err = err
		return false
	}

	parsedEvent, err := es.unmarshaller(e)
	if err != nil {
		es.err = err
		return false
	}

	if publish {
		es.val = &parsedEvent
	} else {
		es.val = nil
	}

	return true
}

// Value returns the most recent event that was generated from a call to Next
func (es *EventStream[T]) Value() *T {
	return es.val
}

// Err returns the first non-EOF error that was encountered
func (es *EventStream[T]) Err() error {
	return es.err
}

// Close will release underlying resources held by an event stream. It must
// always be called.
func (es *EventStream[T]) Close() error {
	return es.r.Close()
}
