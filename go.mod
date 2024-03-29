module github.com/multiverse-os/webframe

go 1.19

require (
	git.mills.io/prologic/bitcask v1.0.2
	github.com/akrylysov/pogreb v0.10.2
	github.com/multiverse-os/muid v0.1.0
	github.com/multiverse-os/service v0.1.2
	golang.org/x/sys v0.15.0
	gopkg.in/yaml.v2 v2.4.0
)

// TODO: And this one gets added in presumably from our example command despite
// our exclude
require github.com/multiverse-os/ansi v0.0.0-20230122075550-10efed87b9d4

exclude (
	github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/banner v0.1.0
)

replace (
	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner => github.com/multiverse-os/banner v0.1.0
)

//replace github.com/multiverse-os/webframe/os/service v0.2.1 => ./os/service
replace github.com/multiverse-os/maglev v0.1.3 => ./app

require (
	github.com/abcum/lcp v0.0.0-20201209214815-7a3f3840be81 // indirect
	github.com/gofrs/flock v0.8.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/plar/go-adaptive-radix-tree v1.0.4 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/exp v0.0.0-20200228211341-fcea875c7e85 // indirect
)
