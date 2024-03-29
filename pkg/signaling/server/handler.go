package server

import (
	"errors"
	"github.com/piotr-gladysz/go-webrtc-tunnel/pkg/signaling/message"
)

var ReceiverNotFoundError = errors.New("Receiver not found")

func (s *SignalingServer) handleSignaling(sess *WSSession, env *message.Envelope) error {
	recvMsg, _ := env.GetSignaling()

	receiver := s.sessions.GetById(recvMsg.Receiver)
	if receiver == nil {

		errMsg := message.NewErrorMessage("Receiver not found", 1, env.Id)
		encodedErr, err := s.encoder.Encode(errMsg)
		if err != nil {
			return err
		}

		if !sess.Send(encodedErr) {
			return errors.New("Failed to send error message")
		}

		return ReceiverNotFoundError
	}

	msg := message.NewSignaling(recvMsg.Data, recvMsg.Receiver, sess.id, env.Type)

	data, err := s.encoder.Encode(msg)
	if err != nil {
		return err
	}

	receiver.Send(data)
	return nil
}
