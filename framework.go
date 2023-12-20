package maglev

import (
	"fmt"
	"strings"
	"time"

	config "github.com/multiverse-os/maglev/config"
	database "github.com/multiverse-os/maglev/database"
	io "github.com/multiverse-os/maglev/io"
	filesystem "github.com/multiverse-os/maglev/os/filesystem"
	router "github.com/multiverse-os/maglev/router"
	server "github.com/multiverse-os/maglev/server"

	ansi "github.com/multiverse-os/ansi"
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
	Config config.Settings

	// TODO: For now we will just have a Debug bool, but this represents later
	// production vs. development paradigm; and we need to have separate
	// initializion to represent this for each system. Proably like a environment
	// pkg that has Init(), and tied into app because app would be where the
	// initialzation happens for both types

	Process     *service.Process
	Directories filesystem.AppDirectories
	Outputs     io.Outputs

	Router *router.Mux

	// TODO: Decide if we prefer naming it Database, and accessing it through
	//           KV(storeType) or Document(storeType) or Columnar(storeType)
	// and if we need to set server.Server to a pointer
	databases map[string]*database.Database
	servers   map[server.Type]server.Server

	Controllers map[string]Controller
	Models      map[string]Model

	// TODO: Jobs should probably be a channel?
	//Jobs []Job

	// TODO: Actually start/stop but several functions need to be renamed to
	// correct this misnaming
	startup  []func()
	shutdown []func()
}

// TODO: id still prefer if this was framework.App not framework.Framework
func Init(cfg config.Settings) Framework {
	service.DropPriviledges()
	service.SeedRandom()

	cfg = config.Validate(cfg)

	framework := Framework{
		Config:      cfg,
		Outputs:     DefaultOutputs(cfg),
		Process:     service.ParseProcess(),
		servers:     make(map[server.Type]server.Server),
		databases:   make(map[string]*database.Database),
		Controllers: make(map[string]Controller),
		Models:      make(map[string]Model),
		// TODO: We should establish some standardized middleware we attach for
		// greater control over our routing infrastructure
		Router: router.New(),
	}

	framework.InitSignalHandler()

	return framework
}

func (f Framework) NewDatabase(name string) {
	name = strings.ToLower(name)
	f.databases[name] = &database.Database{
		Name: name,
	}
}

func (f Framework) NewModel(name string) {
	name = strings.ToLower(name)
	f.Models[name] = Model{
		Name:      name,
		Framework: f,
	}
}

func (f *Framework) NewController(name string) {
	name = strings.ToLower(name)
	f.Controllers[name] = Controller{
		Name:      name,
		Framework: f,
	}
}

func (f Framework) Benchmark(startedAt time.Time, description string) {
	f.Outputs.Write(
		ansi.Bold("Benchmark"),
		ansi.Green(description),
		io.Enclose(ansi.Green(fmt.Sprintf("%v", time.Since(startedAt)))),
	)
}
