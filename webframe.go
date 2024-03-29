package webframe

import (
	"math/rand"
	"strings"
	"time"

	config "github.com/multiverse-os/webframe/config"
	database "github.com/multiverse-os/webframe/database"
	io "github.com/multiverse-os/webframe/io"
	"github.com/multiverse-os/webframe/model"
	filesystem "github.com/multiverse-os/webframe/os/filesystem"
	router "github.com/multiverse-os/webframe/router"
	server "github.com/multiverse-os/webframe/server"

	service "github.com/multiverse-os/service"
)

// TODO: The code overall found in this package is meant to provide access to
// all the various aspects of the framework any application or developer will
// need in order to avoid the need for direct subpackage calls ideally (may not
// be feasible or desirable)

// TODO: Init should NOT be running anything executing anything; functions must
// follow exactly what the function name states; it should only prepare the
// struct.

// As it grows from this original file it should be broken up into logical
// sub-files.
type Framework struct {
	Config *config.Settings

	// TODO: For now we will just have a Debug bool, but this represents later
	// production vs. development paradigm; and we need to have separate
	// initializion to represent this for each system. Proably like a environment
	// pkg that has Init(), and tied into app because app would be where the
	// initialzation happens for both types

	// TODO: I would like to have sub-processes eventually and be able to manage
	// them or maybe it would be better to have a rack like system to segregate
	// the http server then we can put udp or tcp infront of it and pull all
	// that logic out of the framework which would be ideal
	Process     *service.Process
	Directories filesystem.Directories
	Outputs     io.Outputs

	Router *router.Mux

	// TODO: Decide if we prefer naming it Database, and accessing it through
	//           KV(storeType) or Document(storeType) or Columnar(storeType)
	// and if we need to set server.Server to a pointer
	databases map[database.StoreType]*database.Database
	servers   map[server.Type]server.Server

	Controllers map[string]Controller
	Models      map[string]*Model

	// TODO: Jobs should probably be a channel?
	//Jobs []Job

	// TODO: Actually start/stop but several functions need to be renamed to
	// correct this misnaming
	startup  []func()
	shutdown []func()
}

// TODO: id still prefer if this was framework.App not framework.Framework
func Init(cfg *config.Settings) Framework {
	// TODO: We want to always guarantee that no one ever runs a web server as
	// root. This is one of the biggest mistakes a web administrator can make
	// and so we force security.
	service.DropPrivs()
	//service.SeedRandom()
	// TODO: The idea was to put this in service since it should always be done
	// and ideally by process
	rand.Seed(time.Now().UnixNano())

	cfg = config.Validate(cfg)

	framework := Framework{
		Config:      cfg,
		Outputs:     DefaultOutputs(cfg),
		Process:     service.ParseProcess(),
		servers:     make(map[server.Type]server.Server),
		databases:   make(map[database.StoreType]*database.Database),
		Controllers: make(map[string]Controller),
		Models:      make(map[string]*Model),
		// TODO: We should establish some standardized middleware we attach for
		// greater control over our routing infrastructure
		Router: router.New(),
	}

	framework.InitSignalHandler()

	return framework
}

//func (f *Framework) NewDatabase(name string) {
//	name = strings.ToLower(name)
//	f.databases[Model] = &database.Database{
//		Name: name,
//	}
//}

// TODO: Need at least database, well no i guess it has to be the model database
func (f *Framework) NewModel(name string) {
	name = strings.ToLower(name)
	f.Models[name] = &Model{
		Collection: &model.Collection{
			Name: name,
		},
		Framework: f,
	}
}

// TODO: This probably works because we have the logic of controller stored too
//
//	much in the framework but I may ahve a compromise
func (f *Framework) NewController(name string) {
	name = strings.ToLower(name)
	f.Controllers[name] = Controller{
		Name:      name,
		Framework: f,
	}
}
