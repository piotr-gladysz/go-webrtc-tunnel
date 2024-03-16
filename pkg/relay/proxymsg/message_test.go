package proxymsg

import (
	"bytes"
	"testing"
)

func TestProxyDecoder_Decode(t *testing.T) {

	data := make([]byte, 300)

	for i := 0; i < 300; i++ {
		data[i] = byte(i)
	}

	msg := NewProxyMessage(1000, 2000, 3, data)

	enc := &ProxyEncoder{}
	dec := &ProxyDecoder{}

	buf := make([]byte, 1024)

	err := enc.Encode(&msg, buf)
	if err != nil {
		t.Error(err)
	}

	ret, err := dec.Decode(buf)

	if err != nil {
		t.Error(err)
	}

	if ret.Length != msg.Length {
		t.Errorf("Length: %d != %d", msg.Length, ret.Length)
	}

	if ret.Port != msg.Port {
		t.Errorf("Port: %d != %d", msg.Port, ret.Port)
	}

	if ret.ConnId != msg.ConnId {
		t.Errorf("ConnId: %d != %d", msg.ConnId, ret.ConnId)
	}

	if !bytes.Equal(ret.Data, msg.Data) {
		t.Errorf("Data: %v != %v", msg.Data, ret.Data)
	}
}
