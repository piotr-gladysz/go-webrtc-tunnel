package message

import (
	"testing"
)

func TestSimpleJsonDecoder_Decode(t *testing.T) {
	decoder := &SimpleJsonDecoder{}
	encoder := &SimpleJsonEncoder{}

	e := NewErrorMessage("test", 1, 1)
	e.Id = 1

	data, _ := encoder.Encode(e)

	var env Envelope
	err := decoder.Decode(data, &env)

	if err != nil {
		t.Fatalf("expected nil, got %s", err)
	}

	if env.Id != 1 {
		t.Errorf("expected 1, got %d", env.Id)
	}

	if env.Type != MessageTypeError {
		t.Errorf("expected %d, got %d", MessageTypeError, env.Type)
	}

	if env.decodedData == nil {
		t.Errorf("expected not nil, got nil")
	}

	if env.decodedData.(*ErrorMessage).Message != "test" {
		t.Errorf("expected test, got %s", env.decodedData.(*ErrorMessage).Message)
	}
}

func BenchmarkSimpleJsonDecoder_Decode(b *testing.B) {
	decoder := &SimpleJsonDecoder{}
	encoder := &SimpleJsonEncoder{}

	e := NewErrorMessage("test", 1, 1)
	data, _ := encoder.Encode(e)
	for i := 0; i < b.N; i++ {
		var env Envelope
		decoder.Decode(data, &env)
	}

	b.ReportAllocs()
}
