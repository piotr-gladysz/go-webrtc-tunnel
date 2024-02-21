package message

import "errors"

type MessageType int32

var InvalidTypeError = errors.New("invalid message type")

const (
	MessageTypeError MessageType = 0

	MessageTypeAuthToken MessageType = 100
	MessageTypeSDPOffer  MessageType = 101
	MessageTypeSDPAnswer MessageType = 102
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

func (e *Envelope) GetAuthToken() (*AuthToken, error) {
	if e.Type != MessageTypeAuthToken {
		return nil, InvalidTypeError
	}

	ret, ok := e.decodedData.(*AuthToken)
	if !ok {
		return nil, InvalidTypeError
	}

	return ret, nil
}

func (e *Envelope) GetSDP() (*SDP, error) {
	if e.Type != MessageTypeSDPOffer && e.Type != MessageTypeSDPAnswer {
		return nil, InvalidTypeError
	}

	ret, ok := e.decodedData.(*SDP)
	if !ok {
		return nil, InvalidTypeError
	}

	return ret, nil
}
