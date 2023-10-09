package event

import (
	"bytes"
)

// TODO: Segregate Watch code and Event code; then make a library where watch
// uses event and inotify

type Watch struct {
	Name  string    // Relative path to the file or directory.
	Event FileEvent // File operation that triggered the event.
}

// Op (inotify.Op) describes a set of file operations.
type FileEvent uint32

// These are the generalized file operations that can trigger a notification.
const (
	Create FileEvent = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

// String (inotify.Op.String)
func (op Op) String() string {
	// Use a buffer for efficient string concatenation
	var buffer bytes.Buffer

	if op&Create == Create {
		buffer.WriteString("|CREATE")
	}
	if op&Remove == Remove {
		buffer.WriteString("|REMOVE")
	}
	if op&Write == Write {
		buffer.WriteString("|WRITE")
	}
	if op&Rename == Rename {
		buffer.WriteString("|RENAME")
	}
	if op&Chmod == Chmod {
		buffer.WriteString("|CHMOD")
	}
	if buffer.Len() == 0 {
		return ""
	}
	return buffer.String()[1:] // Strip leading pipe
}
