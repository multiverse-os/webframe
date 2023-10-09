package filesystem

import (
	"fmt"
	"io/ioutil"

	color "github.com/multiverse-os/maglev/ansi/color"
)

// NOTE: Provide functionality for printing to terminal, should not be a lot so
// we can condense it all here and only call the ansi library once.

func InitializeFile(path, content string) error {
	fmt.Println("  " + color.Green("CREATE") + " " + path)
	contentBytes := []byte(content)
	return ioutil.WriteFile(path, contentBytes, 0644)
}
