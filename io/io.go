package io

import (
	"fmt"
	"io"
	"os"
	"strings"

	ansi "github.com/multiverse-os/maglev/io/ansi"
)

type Directories struct {
	Working string
	Data    string
	Cache   string
}

const prefixSize = 2
const tabSize = 4

func Prefix() string { return strings.Repeat(" ", prefixSize) }
func Tab() string    { return strings.Repeat(" ", tabSize) }

type Stdio struct {
	In  FileReader
	Out FileWriter
	Err io.Writer
}

type FileWriter interface {
	io.Writer
	Fd() uintptr
}

type FileReader interface {
	io.Reader
	Fd() uintptr
}

type Outputs []Output

// TODO: This provide the formatter logic for extending colors ontop of fmt by
// adding new %X type logic. we can then add %{blue}%{bold} or like css
// %{color:blue;ansi:bold;}
// https://github.com/nhooyr/color
// This is important because its also the founation for a nice implementaiton of
// locales without relying on outside depndencies

// The hex output for this formatter is nice, rest is ok
// https://github.com/go-ffmt/ffmt

// A lot of logic should be handled by either terminal or text libraries (and
// these may be merged again)
// https://github.com/jedib0t/go-pretty
// Text
//
// Utility functions to manipulate text with or without ANSI escape sequences. Most of the functions available are used in one or more of the other packages here.
//     Align text horizontally or vertically
//         text/align.go and text/valign.go
//     Colorize text
//         text/color.go
//     Cursor Movement
//         text/cursor.go
//     Format text (convert case)
//         text/format.go
//     String Manipulation (Pad, RepeatAndTrim, RuneCount, Trim, etc.)
//         text/string.go
//     Transform text (UnixTime to human-readable-time, pretty-JSON, etc.)
//         text/transformer.go
//     Wrap text
//         text/wrap.go

// TODO: Formatters like this will provide the best way forward for not just
// doing ANSI better, but also provide simple way to handle localizaiton.
// https://github.com/kr/pretty/blob/main/formatter.go

// TODO: Consider rebuilding the table library, do it in such a way that each of
// the visual compoenents are mapped to a type that can be loaded like a theme.
// https://github.com/konojunya/go-frame/blob/master/frame.go

// The same should be with tree structures. Also the tree structure should
// support horizontal and vertical printing.

// THe concept of wrapping at 80 is great and would simplify our help output code
// https://github.com/PraserX/afmt

// TODO: Use terminal library to read width and add ability to format to one side
// Logging that ideally is not too bloated to get in the way, support for
// overriding by passing your logger's os.Writer but enough complexity to be
// useful in many use cases.

type Output struct {
	//Timestamp string
	prefix    string
	stripANSI bool
	file      io.Writer
}

// TODO: When this goes into its own package, this should be moved to its own file
type LogLevel int

const (
	LOG LogLevel = iota
	INFO
	DEBUG
	WARN
	ERROR
	FATAL
)

func (self LogLevel) String() string {
	switch self {
	case INFO:
		return ansi.SkyBlue("info")
	case DEBUG:
		return ansi.Blue("debug")
	case WARN:
		return ansi.Olive("warn")
	case ERROR:
		return ansi.Red("error")
	case FATAL:
		return ansi.Maroon("fatal")
	default: // LOG
		return ansi.Light("log")
	}
}

func merge(textParts ...interface{}) (output string) {
	for _, part := range textParts {
		output += fmt.Sprint(part, " ")
	}
	return output
}

// Value Assignment Chaining //////////////////////////////////////////////////
func (self Output) Prefix(prefix string) Output {
	self.prefix = prefix
	return self
}

// TODO: By using the method of coloring that does it via % using fmt override.
// WE can just not place the ANSI in. Since filelogging is more likely usecase
// it will use less resources in the long run to do that approach over the strip
// approach
func (self Output) StripANSI(strip bool) Output {
	self.stripANSI = strip
	return self
}

// Default Ouput Locations ////////////////////////////////////////////////////
func TerminalOutput() Output {
	//palette, err := colorful.WarmPalette(10)
	//if err != nil {
	//	return nil, err
	//}
	return Output{
		//Theme: Theme{
		//	//Primary:   palette[0],
		//	//Secondary: palette[1],
		//	//Contrast:  palette[2],
		//},
		prefix:    "starshipyard",
		file:      os.Stdout,
		stripANSI: false,
	}
}

func LogfileOutput(filename string) Output {
	// TODO: Create if does not exist, including path. For now, lets just declare
	// log rotation as out of scope. If its not widely supported you require a
	// third party program to handle it anyways, so just leave it to that.
	return Output{
		stripANSI: true,
	}
}

// Writing To Outputs /////////////////////////////////////////////////////////
// TODO: Improve this by implmenting Fprintf locally, so we can provide similar
// functionality to Ouput and Write.
func (self Outputs) Write(text ...interface{}) {
	for _, output := range self {
		output.Write(text)
	}
}

func (self Output) Write(text ...interface{}) {
	var output string
	for _, t := range text {
		output += fmt.Sprint(t)
	}

	if self.stripANSI {
		fmt.Fprintf(self.file, "%s", self.prefix+ansi.Regex.ReplaceAllString(fmt.Sprint(text), "")+"\n")
	} else {
		//fmt.Fprint(self.file, self.prefix, fmt.Sprint(text), "\n")
		fmt.Fprint(self.file, fmt.Sprintf("%s%v\n", Enclose(self.prefix), output))
	}
}

func VarName(text string) string {
	return ansi.Bold(ansi.Blue(text))
}

func Enclose(text string) string {
	return ansi.Bold(ansi.White("[")) + text + ansi.Bold(ansi.White("]"))
}

func (self Output) Log(level LogLevel, text ...interface{}) {
	var output string
	for _, t := range text {
		output += fmt.Sprintf("%v", t)
	}
	self.Write(Enclose(level.String()), " ", output)
}

// TODO: [[!!]] Functions for ansi access so we dont need to include it elsewhere and
// alter implement it locally here so we haev no deps

// TODO: Would be much better to just allow providing a closure, or a
// surrounding fucntion, whatever that function does. And in our case it will be
// coloring. That way its more flexible, needs less color specific code, and
// allows developers to pick colors.
func (self Outputs) Log(level LogLevel, text ...interface{}) {
	for _, output := range self {
		output.Log(level, text...)
	}
}
