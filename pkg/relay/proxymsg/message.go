package proxymsg

import "errors"

var InvalidLengthError = errors.New("invalid message length")

const (
	ActionData  = uint8(1)
	ActionClose = uint8(2)
	ActionError = uint8(3)
)

type ProxyMessage struct {
	Length uint16
	Port   uint16
	ConnId uint16
	Action uint8
	Data   []byte
}

func NewProxyMessage(port, connId uint16, action uint8, data []byte) ProxyMessage {
	return ProxyMessage{
		Length: uint16(len(data)),
		Port:   port,
		ConnId: connId,
		Action: action,
		Data:   data,
	}
}

func (p *ProxyMessage) GetTotalLen() int {
	return int(p.Length) + 7
}

type ProxyDecoder struct {
}

func (p *ProxyDecoder) Decode(data []byte) (ProxyMessage, error) {

	var ret ProxyMessage

	ret.Length = uint16(data[0]) | uint16(data[1])<<8

	if len(data) < int(ret.Length)+7 {
		return ret, InvalidLengthError
	}

	ret.Port = uint16(data[2]) | uint16(data[3])<<8
	ret.ConnId = uint16(data[4]) | uint16(data[5])<<8
	ret.Action = data[6]
	ret.Data = data[7 : ret.Length+7]

	return ret, nil
}

type ProxyEncoder struct {
}

func (p *ProxyEncoder) Encode(msg *ProxyMessage, buf []byte) error {

	buf[0] = byte(msg.Length)
	buf[1] = byte(msg.Length >> 8)

	buf[2] = byte(msg.Port)
	buf[3] = byte(msg.Port >> 8)

	buf[4] = byte(msg.ConnId)
	buf[5] = byte(msg.ConnId >> 8)

	buf[6] = msg.Action

	copy(buf[7:], msg.Data)

	return nil
}
