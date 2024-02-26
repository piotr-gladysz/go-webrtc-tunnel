package message

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

// SimpleJsonDecoder is a simple decoder that decodes messages from JSON.
// To be optimized later.
type SimpleJsonDecoder struct {
}

func (s *SimpleJsonDecoder) Decode(data []byte, msg *Envelope) error {

	err := json.Unmarshal(data, &msg)

	if err != nil {
		return err
	}

	switch msg.Type {
	case MessageTypeError:
		body := new(ErrorMessage)

		err := mapstructure.Decode(msg.Data, body)
		if err != nil {
			return err
		}

		msg.decodedData = body
	case MessageTypeAuthToken:
		body := new(AuthInfo)

		err := mapstructure.Decode(msg.Data, body)
		if err != nil {
			return err
		}

		msg.decodedData = body
	case MessageTypeSDPOffer, MessageTypeSDPAnswer:
		body := new(SDP)

		err := mapstructure.Decode(msg.Data, body)
		if err != nil {
			return err
		}

		msg.decodedData = body
	}

	return nil

}
