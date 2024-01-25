package main

import (
	"fmt"
	"os"

	app "github.com/multiverse-os/maglev-app"

	cli "github.com/multiverse-os/cli"
	config "github.com/multiverse-os/maglev/config"
)

// TODO: Review the additional functionality provided by Rails binary, so that
// features like `rails notes` can be added (which scans files for TODO and for
// our purposes NOTE, and likely define what it looks for via YAML configuration
// for a generally useful system; then take these TODOs and others and build a
// notes file to help guide development). [not neccessarily that, or at least
// not necessarily high priority, many of these features that are applicable
// should be added
func main() {
	cmd, _ := cli.New(cli.App{
		Name: "maglev",
		// TODO: Consider method from config giving us the outputs to avoid the
		// extra import (for simplicity not because it actually reduces resources
		// beacuse it naturally is eventually called in by app pkg
		// Doesn't work but something like this so we can provide the cli framework
		// outputs defined by our higher level application [consistency!]
		//Outputs:     cli.Outputs{framework.DefaultOutputs()},
		Version:     cli.Version{Major: 0, Minor: 1, Patch: 1},
		Description: "A command-line tool for controling the maglev server, scaffolding boilerplate code, and executing developer defined commands",
		GlobalFlags: cli.Flags(
			cli.Flag{
				Category:    "Server",
				Name:        "env",
				Alias:       "e",
				Default:     "development",
				Description: "Specify the server environment",
			},
			cli.Flag{
				Category:    "Server",
				Name:        "address",
				Alias:       "a",
				Default:     "0.0.0.0",
				Description: "Specify the address for the HTTP server to listen",
			},
			cli.Flag{
				Category:    "Server",
				Name:        "port",
				Alias:       "p",
				Default:     "3000",
				Description: "Specify the listening port for the HTTP server",
			},
		),
		Commands: cli.Commands(
			cli.Command{
				Name:        "server",
				Alias:       "s",
				Description: "Options for controlling maglev HTTP server",
				Subcommands: cli.Commands(
					cli.Command{
						Name:        "start",
						Alias:       "s",
						Description: "Start the maglev http server",
						Flags: cli.Flags(
							cli.Flag{
								Name:        "daemonize",
								Alias:       "d",
								Description: "Daemonize the http server",
							},
						),
						Action: func(c *cli.Context) error {
							cfg, err := config.Load("app/config/app.yaml")
							if err != nil {
								cfg = config.DefaultConfig(cfg.Name)
							}

							// TODO: We also need helpers for gathering environmental variabls
							// to cascade through: file config => cli config => env config
							// with env config overriding all other methods of configuration
							cfg.Address = c.Flag("address").String()
							cfg.Port = c.Flag("port").Int()

							// TODO: If inside the CLI application we are using server.Log
							// here, we need to be consistent; we should initialize an
							// app.App object and assign the default Outputs to cli.Outputs
							server := app.Init(cfg)
							if c.Flag("daemonize").Bool() {
								server.Log("launching in daemon mode...")
								server.Log("not implemented")
							} else {
								server.Log("launching with terminal attached to server")
								server.Log("http server listening on [", cfg.Address, ":", cfg.Port, "]")
								server.Start()
							}
							return nil
						},
					},
				),
			},
			cli.Command{
				Name:        "generate",
				Alias:       "g",
				Description: "Generate new go source code for models, controllers, and views",
				Subcommands: cli.Commands(
					cli.Command{
						Name:        "model",
						Description: "Build a model template with the specified object data",
						Action: func(c *cli.Context) error {
							fmt.Println("code generation functionality is not implemented yet")
							fmt.Println("test code has been built, it just needs to be migrated into the base and will be available shortly")
							return nil
						},
					},
					cli.Command{
						Name:        "controller",
						Description: "Build a controller template with the specified object data",
						Action: func(c *cli.Context) error {
							fmt.Println("code generation functionality is not implemented yet")
							fmt.Println("test code has been built, it just needs to be migrated into the base and will be available shortly")
							return nil
						},
					},
					cli.Command{
						Name:        "view",
						Description: "Build a view template with the specified object data",
						Action: func(c *cli.Context) error {
							fmt.Println("code generation functionality is not implemented yet")
							fmt.Println("test code has been built, it just needs to be migrated into the base and will be available shortly")
							return nil
						},
					},
					cli.Command{
						Name:        "job",
						Description: "Build a job template with the specified object data",
						Action: func(c *cli.Context) error {
							fmt.Println("code generation functionality is not implemented yet")
							fmt.Println("test code has been built, it just needs to be migrated into the base and will be available shortly")
							return nil
						},
					},
				),
			},
			cli.Command{
				Name:        "console",
				Alias:       "c",
				Description: "Start the maglev yard console interface",
				Action: func(c *cli.Context) error {
					fmt.Println("[CONSOLE] console interface is not implemented yes")
					return nil
				},
			}),
	},
	)

	cmd.Parse(os.Args).Execute()
}
