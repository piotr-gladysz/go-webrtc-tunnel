package message

// messageId is a global variable that is used to assign unique IDs to messages.
var messageId int32 = 0

type ErrorMessage struct {
	Message   string `json:"message"`
	Code      int32  `json:"code"`
	RelatedId int32  `json:"related_id"`
}

type AuthInfo struct {
	Token string `json:"token"`
}

type Signaling struct {
	Data     string `json:"data"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
}

func NewErrorMessage(message string, code int32, relatedId int32) *Envelope {
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
		decodedData: &AuthInfo{
			Token: token,
		},
	}
}

func NewSignaling(sdp, receiver, sender string, msgType MessageType) *Envelope {
	messageId++
	return &Envelope{
		Id:   messageId,
		Type: msgType,
		decodedData: &Signaling{
			Data:     sdp,
			Receiver: receiver,
			Sender:   sender,
		},
	}
}
