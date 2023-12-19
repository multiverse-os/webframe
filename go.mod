module github.com/multiverse-os/maglev

go 1.19

replace (
	github.com/multiverse-os/maglev/database => ./database
	github.com/multiverse-os/maglev/os/service => ./os/service
)

require github.com/multiverse-os/maglev/os/service v0.1.0

require github.com/multiverse-os/maglev/database v0.1.0
