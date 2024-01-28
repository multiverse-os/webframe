package webframe

import (
	server "github.com/multiverse-os/webframe/server"
)

func (f *Framework) HTTP() server.Server {
	var err error
	if f.servers[server.HTTP] == nil {
		f.servers[server.HTTP], err = server.New(server.HTTP, f.Config.Address, f.Config.Port)
		if err != nil {
			panic(err.Error())
		}
	}
	return f.servers[server.HTTP]
}
