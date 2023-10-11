package service

import (
	"fmt"
	"os"
	"reflect"
	"syscall"
	"time"
	"unsafe"

	pid "github.com/multiverse-os/service/pid"
	signal "github.com/multiverse-os/service/signal"
)

// [ General process utilities ]////////////////////////////////////////////////
func isProcessRunning(pid int) bool {
	if process, err := os.FindProcess(pid); err == nil {
		return false
	} else {
		err = process.Signal(syscall.Signal(0))
		return (err != nil)
	}
}

//func SetUserId(procAttr *syscall.SysProcAttr, uid uint32, gid uint32) {
//	procAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid, NoSetGroups: true}
//}

// TODO: May want to add a weight concept so that certain functions will be
// esnured to be ran before others
type Process struct {
	Pid                int
	PidFile            *pid.File
	User               string
	UID                int
	GID                int
	TempDirectory      string
	UserCacheDirectory string
	Env                map[string]string
	Executable         string
	WorkingDirectory   string
	StartedAt          time.Time
	Signals            signal.Handler
	shutdown           []func()
}

// TODO
// This started at isn't the actual started at time, its the parsed at
// time.
func ParseProcess() *Process {
	if executable, err := os.Executable(); err != nil {
		panic(fmt.Errorf("failed to parse process:", err))
	} else {
		return &Process{
			Pid:        os.Getpid(),
			UID:        os.Getuid(),
			GID:        os.Getgid(),
			Executable: executable,
			StartedAt:  time.Now(),
		}
	}
}

func (p *Process) ShutdownCount() int { return len(p.shutdown) }

func (p *Process) Shutdown() {
	for _, function := range p.shutdown {
		function()
	}
}

func (p *Process) AppendToShutdown(exitFunction func()) {
	p.shutdown = append(p.shutdown, exitFunction)
}

// [ Method for process ]///////////////////////////////////////////////////////
// NOTE: Returns the close function so that it can be called easily added to a
// defer. This is important because since we are doing OS based locks on the
// pidfile we may need to unlock the file
func (p *Process) WritePid(path string) *pid.File {
	fmt.Println("writing pid:", p.Pid)
	if pidFile, err := pid.Write(path); err != nil {
		panic(fmt.Errorf("failed to write pid: ", err))
	} else {
		p.PidFile = pidFile
		return p.PidFile
	}
}

func (p *Process) CleanPid() error { return p.PidFile.Clean() }

// NOTE: Set process name, as in the name seen in `ps`
func (self *Process) SetName(name string) {
	argv0str := (*reflect.StringHeader)(unsafe.Pointer(&os.Args[0]))
	argv0 := (*[1 << 30]byte)(unsafe.Pointer(argv0str.Data))[:argv0str.Len]
	n := copy(argv0, name)
	if n < len(argv0) {
		argv0[n] = 0
	}
}
