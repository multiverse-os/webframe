package daemon

import "errors"

var (
	errStop         = errors.New("stop serve signals")
	errWouldBlock   = errors.New("resource temporarily unavailable")
	errNotSupported = errors.New("Non-POSIX OS is not supported")
)
