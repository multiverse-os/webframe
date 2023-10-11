package server

import "errors"

var (
	errServerWS  = errors.New("defined server type ws is currently unavailable")
	errServerRTC = errors.New("defined server type rtc is currently unavailable")
	errServerTCP = errors.New("defined server type tcp is currently unavailable")
	errServerUDP = errors.New("defined server type udp is currently unavailable")
)
