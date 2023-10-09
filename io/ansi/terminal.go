package ansi

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type WindowSize struct {
	rows    uint16
	columns uint16
	width   uint16 // Pixel
	height  uint16 // Pixel
}

// TODO: Interesting with the bufio.Writer and Reader, this could be
// foundational for the IO model we are building in the application
// skeleton (starting to develop into a MVOS go framework)
type Terminal struct {
	cursor *Cursor
	width  int
	height int
	stdout bufio.Writer
	stdin  bufio.Reader
}

// TODO: Is 64 enough?
func NewTerminal() Terminal {
	terminal := Terminal{
		width:  16,
		height: 16,
		stdout: *bufio.NewWriterSize(os.Stdout, 4096),
		stdin:  *bufio.NewReaderSize(os.Stdin, 64),
	}
	terminal.UpdateSize()
	return terminal
}

// Update our size to match the real TTY session
func (self *Terminal) UpdateSize() {
	var windowSize WindowSize
	syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(&windowSize)),
	)

	self.width = int(windowSize.columns)
	self.height = int(windowSize.rows)
}

// NOTE: This is very fucking nice
func (self *Terminal) Draw(callback func()) {
	callback()
	self.Flush()
}

func (self *Terminal) Size() (int, int) { return self.width, self.height }
func (self *Terminal) Width() int       { return self.width }
func (self *Terminal) Height() int      { return self.height }
func (self *Terminal) Flush()           { self.stdout.Flush() }

// NOTE: This is truly very nice
func (self *Terminal) Write(format string, a ...interface{}) { fmt.Fprintf(&self.stdout, format, a...) }

func (self *Terminal) WriteRune(char rune) {
	fmt.Fprint(&self.stdout, "%s", string(char))
}
