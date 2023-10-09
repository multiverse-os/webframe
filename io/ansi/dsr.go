package ansi

import "fmt"

func (self *Terminal) DSR(a ...interface{}) {
	fmt.Fprintf(&self.stdout, "\033[", a...) // byte(27)
}
