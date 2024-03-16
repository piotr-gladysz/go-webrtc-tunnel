package message

import "errors"

type MessageType int32

var InvalidTypeError = errors.New("invalid message type")

const (
	MessageTypeError MessageType = 0

	MessageTypeAuthToken MessageType = 100

	MessageTypeSDPOffer     MessageType = 101
	MessageTypeSDPAnswer    MessageType = 102
	MessageTypeICECandidate MessageType = 103
)

type Envelope struct {
	Id   int32          `json:"id"`
	Type MessageType    `json:"type"`
	Data map[string]any `json:"data"`

	decodedData any
}

func (e *Envelope) GetError() (*ErrorMessage, error) {
	if e.Type != MessageTypeError {
		return nil, InvalidTypeError
	}

	ret, ok := e.decodedData.(*ErrorMessage)
	if !ok {
		return nil, InvalidTypeError
	}

	return ret, nil
}

func (e *Envelope) GetAuthToken() (*AuthInfo, error) {
	if e.Type != MessageTypeAuthToken {
		return nil, InvalidTypeError
	}

	ret, ok := e.decodedData.(*AuthInfo)
	if !ok {
		return nil, InvalidTypeError
	}

	return ret, nil
}

func (e *Envelope) GetSignaling() (*Signaling, error) {
	if e.Type != MessageTypeSDPOffer && e.Type != MessageTypeSDPAnswer && e.Type != MessageTypeICECandidate {
		return nil, InvalidTypeError
	}

	ret, ok := e.decodedData.(*Signaling)
	if !ok {
		return nil, InvalidTypeError
	}

	return ret, nil
}
