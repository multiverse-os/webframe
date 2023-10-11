package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	router "github.com/multiverse-os/maglev/router"
)

// TODO: I really love this segregated router from server and just using the
// http server from standard library; leaves us open for switching out the
// different components easily

// TODO: Should server object have the app object address? -- well then we could
// output all of the server stuff to the correct output (weather that fd be
// stdout or a log file or both) ---BUT we cant have framework.App beacuse that
// would be a circular call.
// So overtime we will need each of the section of framework to be in their own
// exposed module standardized api to app which runs the process, but also some
// other app which ties the modules together and makes them usable from a single
// object that can be passed trhough without circular calls
// because again for example, being able to write or read from config, being
// able to write files to ~/.local/share/app-name or config or logs will all be
// needed from ALL the submodules (example server, databsae, model)
type Server struct {
	Config Config
	Render http.ResponseWriter
	Router router.Router
	HTTP   *http.Server
	TLS    *TLS
}

func NewServer(address string, port int) *Server {
	return &Server{
		Router: router.New(),
		Config: Config{
			Address: address,
			Port:    port,
		},
	}
}

func (self *Server) UseRouter(r router.Router) {
	self.Router = r
}

// TODO: phase this out
func (self *Server) Type() string    { return "HTTP" }
func (self *Server) IsRunning() bool { return self.HTTP != nil }

func (self *Server) Start() error {
	// TODO: SErver needs access to SOME app or another! so we can do outputs to
	// either log or stdout or whatever the fds are
	self.HTTP = &http.Server{Addr: self.ListeningAt(), Handler: self.Router}
	self.HTTP.ListenAndServe()
	return nil
}

func (self *Server) Stop() error {
	ctx, _ := context.WithTimeout(context.Background(), (15 * time.Second))
	if err := self.HTTP.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown the http server:", err)
	}
	return nil
}

func (self *Server) Restart() (err error) {
	err = self.Stop()
	err = self.Start()
	return err
}

func (self Server) ListeningAt() string {
	return self.Config.Address + ":" + strconv.Itoa(self.Config.Port)
}
