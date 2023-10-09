package http

import (
	"net/http"
)

type TLS struct {
	*http.Server
	CertFile string
	KeyFile  string
}
