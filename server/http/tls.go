package http

import (
	"net/http"
)

// TODO: Should probably have something else do termination
type TLS struct {
	*http.Server
	CertificateFile string
	KeyFile         string
}
