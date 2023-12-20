module github.com/multiverse-os/maglev

go 1.19

require (
	github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/muid v0.1.0
	github.com/multiverse-os/service v0.1.1
	golang.org/x/sys v0.15.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.1.0

replace (
	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.1.0
	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner => github.com/multiverse-os/banner v0.1.0
)

//replace github.com/multiverse-os/maglev-app v0.1.1 => ./app

require golang.org/x/crypto v0.17.0 // indirect
