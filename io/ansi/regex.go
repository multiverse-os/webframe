package ansi

import (
	"regexp"
)

var Regex = regexp.MustCompile("\x1b[@-_][0-?]*[ -/]*[@-~]")
