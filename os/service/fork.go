package service

import (
	"os"
	"os/exec"
	"reflect"
	"syscall"
	"unsafe"
)

func IsProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return (err == nil)
}

func Fork() (int, error) {
	args := os.Args[1:]
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = os.Environ()
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.ExtraFiles = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Setsid is used to detach the process from the parent
		// (normally a shell)
		//
		// The disowning of a child process is
		// accomplished by executing the system call
		// setpgrp() or setsid(), (both of which
		// have the same functionality) as soon as
		// the child is forked. These calls
		// create a new process session group,
		// make the
		// child process the session leader,
		// and set the process group ID to
		// the process
		// ID of the child.
		// https://bsdmag.org/unix-kernel-system-calls/
		Setsid: true,
	}
	if err := cmd.Start(); err != nil {
		return 0, err
	}
	return cmd.Process.Pid, nil
}

// NOTE: Set process name, as in the name seen in `ps`
func setProcessName(name string) {
	argv0str := (*reflect.StringHeader)(unsafe.Pointer(&os.Args[0]))
	argv0 := (*[1 << 30]byte)(unsafe.Pointer(argv0str.Data))[:argv0str.Len]
	n := copy(argv0, name)
	if n < len(argv0) {
		argv0[n] = 0
	}
}

func SetUserId(procAttr *syscall.SysProcAttr, uid uint32, gid uint32) {
	procAttr.Credential = &syscall.Credential{Uid: uid, Gid: gid, NoSetGroups: true}
}
