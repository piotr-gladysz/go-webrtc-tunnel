package peer

import (
	"context"
	"errors"
	"github.com/pion/webrtc/v4"
	"sync"
)

const mainDataChannelLabel = "control"

var AlreadyStartedError = errors.New("tunnel already started")

type WebRTCListener interface {
	OnICECandidate(p *Peer, candidate *webrtc.ICECandidate)
	OnLocalDescription(p *Peer, sdp webrtc.SessionDescription)
	OnConnected(p *Peer)
}

type StateListener interface {
	OnICEConnectionStateChange(p *Peer, state webrtc.ICEConnectionState)
	OnDataChannel(p *Peer, channel *webrtc.DataChannel)
}

type Peer struct {
	Id                string
	AllowedLocalPorts []int
	AllowedAll        bool

	parentCtx context.Context

	ctx    context.Context
	cancel context.CancelFunc

	stunServers []string

	isStarted bool

	lastError   error
	isConnected bool
	connCh      chan struct{}

	peerConnection *webrtc.PeerConnection

	iceCandidates []string
	candidateMux  sync.Mutex

	webrtclistener WebRTCListener
	stateListener  StateListener

	controlChannel *webrtc.DataChannel
}

func NewTunnel(parentCtx context.Context, id string, stunServers []string, listener WebRTCListener) *Peer {
	return &Peer{
		Id:             id,
		parentCtx:      parentCtx,
		stunServers:    stunServers,
		webrtclistener: listener,
	}
}

func (p *Peer) Start(descrStr string) (err error) {
	if p.isStarted {
		return AlreadyStartedError
	}

	p.ctx, p.cancel = context.WithCancel(p.parentCtx)

	peerConnection, err := webrtc.NewPeerConnection(p.webrtcConfig())
	if err != nil {
		return err
	}

	p.isStarted = true
	p.peerConnection = peerConnection

	defer func() {
		if err != nil {
			peerConnection.Close()
			p.peerConnection = nil
			p.controlChannel = nil
		}

	}()

	peerConnection.OnICECandidate(p.onICECandidate)
	peerConnection.OnICEConnectionStateChange(p.onICEConnectionStateChange)
	peerConnection.OnDataChannel(p.onDataChannel)
	if p.controlChannel, err = peerConnection.CreateDataChannel(mainDataChannelLabel, nil); err != nil {
		return err
	}

	var localDescription webrtc.SessionDescription
	if descrStr != "" {
		descr := webrtc.SessionDescription{
			Type: webrtc.SDPTypeOffer,
			SDP:  descrStr,
		}
		if err = peerConnection.SetRemoteDescription(descr); err != nil {
			return err
		}
		localDescription, err = peerConnection.CreateAnswer(nil)

		if err != nil {
			return err
		}
	} else {
		localDescription, err = peerConnection.CreateOffer(nil)
		if err != nil {
			return
		}
	}

	if err = peerConnection.SetLocalDescription(localDescription); err != nil {
		return err
	}

	p.webrtclistener.OnLocalDescription(p, localDescription)

	return nil
}

func (p *Peer) Close() error {
	if p.peerConnection == nil {
		return nil
	}

	err := p.peerConnection.Close()
	p.peerConnection = nil
	p.controlChannel = nil

	return err
}

func (p *Peer) SetStateListener(listener StateListener) {
	p.stateListener = listener
}

func (p *Peer) webrtcConfig() webrtc.Configuration {
	return webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: p.stunServers,
			},
		},
	}
}
