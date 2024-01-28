package webframe

import (
	"fmt"
	"time"

	ansi "github.com/multiverse-os/ansi"
	io "github.com/multiverse-os/webframe/io"
)

func (f Framework) Benchmark(startedAt time.Time, description string) {
	f.Outputs.Write(
		ansi.Bold("Benchmark"),
		ansi.Green(description),
		io.Enclose(ansi.Green(fmt.Sprintf("%v", time.Since(startedAt)))),
	)
}
