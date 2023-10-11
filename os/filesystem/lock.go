package filesystem

import (
	"context"
	"os"
	"syscall"
)

func Lock(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
}

func TryLock(file *os.File) error {
	err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	switch err {
	case syscall.EWOULDBLOCK:
		return errIsLocked
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
		return errTimeout
	}
}

func Unlock(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}
