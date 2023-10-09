package filesystem

import (
	"context"
	"errors"
	"os"
	"syscall"
)

var (
	ErrIsLocked = errors.New("file is already locked")
	ErrTimeout  = errors.New("timeout exceeded when acquiring lock")
)

func Lock(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
}

func TryLock(f *os.File) error {
	err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	switch err {
	case syscall.EWOULDBLOCK:
		return ErrIsLocked
	default:
		return err
	}
}

func LockWithContext(ctx context.Context, file *os.File) error {
	result := make(chan error)
	go func() {
		result <- Lock(file)
	}()

	select {
	case err := <-result:
		return err
	case <-ctx.Done():
		go Unlock(file)
		return ErrTimeout
	}
}

func Unlock(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}
