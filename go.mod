module github.com/multiverse-os/maglev

go 1.19

require (
	github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/color v0.1.0
	github.com/multiverse-os/maglev/os/service v0.1.0
	github.com/multiverse-os/pid v0.1.0
	github.com/multiverse-os/signal v0.1.0
)

replace github.com/multiverse-os/service => github.com/multiverse-os/maglev/os/service v0.1.0

replace github.com/multiverse-os/maglev/os/service => ./os/service
