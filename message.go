package wsqueue

import (
	"encoding/json"
	"strconv"

	"github.com/satori/go.uuid"
)

type Header map[string]string

//Message message
type Message struct {
	Header Header `json:"metadata,omitempty"`
	Body   string `json:"data"`
}

func newMessage(data interface{}) (*Message, error) {
	m := Message{
		Header: make(map[string]string),
		Body:   "",
	}

	switch data.(type) {
	case string, *string:
		m.Body = data.(string)
	case int, *int, int32, *int32, int64, *int64:
		m.Body = strconv.Itoa(data.(int))
	case bool, *bool:
		m.Body = strconv.FormatBool(data.(bool))
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		m.Body = string(b)
	}
	m.Header["id"] = uuid.NewV1().String()
	return &m, nil
}

func (m *Message) String() string {
	var s string
	s = "\n---HEADER---"
	for k, v := range m.Header {
		s = s + "\n" + k + ":" + v
	}
	s = s + "\n---BODY---"
	s = s + "\n" + m.Body
	return s
}

//ID returns message if
func (m *Message) ID() string {
	return m.Header["id"]
}