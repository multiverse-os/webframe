package server

import (
	"errors"

	router "github.com/multiverse-os/webkit/router"
	http "github.com/multiverse-os/webkit/server/http"
)

type Type uint

const (
	Undefined Type = iota
	HTTP
	WS
	WebRTC
	TCP
	UDP
)

type Server interface {
	Start() error
	Stop() error
	Restart() error
	IsRunning() bool
	ListeningAt() string
	Type() string

	UseRouter(router.Router)
}

type Servers []Server

// TODO: Then do stop start restart on the collection. So we can do
// Servers.Restart() and it restarts all the servers conviently

func New(serverType Type, address string, port int) (Server, error) {
	switch serverType {
	case HTTP:
		return http.NewServer(address, port), nil
	case WS:
		return nil, errors.New("defined server type ws is currently unavailable")
	case WebRTC:
		return nil, errors.New("defined server type webrtc is currently unavailable")
	case TCP:
		return nil, errors.New("defined server type tcp is currently unavailable")
	case UDP:
		return nil, errors.New("defined server type udp is currently unavailable")
	default: // Undefined
		return nil, errors.New("undefined server type")
	}
}
