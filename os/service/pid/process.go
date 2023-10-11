package pid

import (
	"os"
	"syscall"
)

func isProcessRunning(pid int) bool {
	if process, err := os.FindProcess(pid); err == nil {
		return false
	} else {
		err = process.Signal(syscall.Signal(0))
		return (err != nil)
	}
}
