package message

// messageId is a global variable that is used to assign unique IDs to messages.
var messageId int32 = 0

type ErrorMessage struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	RelatedId int32  `json:"related_id"`
}

type AuthToken struct {
	Token string `json:"token"`
}

type SDP struct {
	SDP      string `json:"sdp"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
}

func NewErrorMessage(message string, code int, relatedId int32) *Envelope {
	messageId++
	return &Envelope{
		Id:   messageId,
		Type: MessageTypeError,
		decodedData: &ErrorMessage{
			Message:   message,
			Code:      code,
			RelatedId: relatedId,
		},
	}
}

func NewAuthToken(token string) *Envelope {
	messageId++
	return &Envelope{
		Id:   messageId,
		Type: MessageTypeAuthToken,
		decodedData: &AuthToken{
			Token: token,
		},
	}
}

func NewSDP(sdp, receiver, sender string, msgType MessageType) *Envelope {
	messageId++
	return &Envelope{
		Id:   messageId,
		Type: msgType,
		decodedData: &SDP{
			SDP:      sdp,
			Receiver: receiver,
			Sender:   sender,
		},
	}
}
