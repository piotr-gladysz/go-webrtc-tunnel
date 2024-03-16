package message

import (
	"encoding/json"
)

// SimpleJsonEncoder is a simple encoder that encodes messages into JSON.
// To be optimized later.
type SimpleJsonEncoder struct {
}

func (s SimpleJsonEncoder) Encode(msg *Envelope) ([]byte, error) {
	data := map[string]interface{}{
		"id":   msg.Id,
		"type": msg.Type,
		"data": msg.decodedData,
	}

	return json.Marshal(data)
}
