module github.com/multiverse-os/service

go 1.19

replace (
	github.com/multiverse-os/service/daemon => ./daemon
	github.com/multiverse-os/service/pid => ./pid
	github.com/multiverse-os/service/signal => ./signal
)
