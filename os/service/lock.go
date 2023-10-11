package service

import (
	"syscall"
)

func lockFile(fd uintptr) error   { return syscall.Flock(int(fd), syscall.LOCK_EX|syscall.LOCK_NB) }
func unlockFile(fd uintptr) error { return syscall.Flock(int(fd), syscall.LOCK_UN) }

//type Locker interface {
//	LockRead() error
//	LockWrite() error
//	WaitAndLockRead() error
//	WaitAndLockWrite() error
//	Unlock()
//}
//
//type Lockfile struct {
//	Path         string
//	file         *os.File
//	maintainFile bool
//	ft           *syscall.Flock_t
//}
//
//func NewLockfile(path string) *Lockfile {
//	return &Lockfile{Path: path, maintainFile: true}
//}
//
//func NewLockfileFromFile(file *os.File) *Lockfile {
//	return &Lockfile{file: file, maintainFile: false}
//}
//
//func (self *Lockfile) LockRead() error {
//	return self.lock(false, false, 0, os.SEEK_SET, 0)
//}
//
//func (self *Lockfile) LockWrite() error {
//	return self.lock(true, false, 0, os.SEEK_SET, 0)
//}
//
//func (self *Lockfile) WaitAndLockRead() error {
//	return self.lock(false, true, 0, os.SEEK_SET, 0)
//}
//
//func (self *Lockfile) WaitAndLockWrite() error {
//	return self.lock(true, true, 0, os.SEEK_SET, 0)
//}
//
//func (self *Lockfile) Unlock() {
//	self.unlock(0, os.SEEK_SET, 0)
//}
//
//func (self *Lockfile) LockReadRange(offset int64, to int, len int64) error {
//	return self.lock(false, false, offset, to, len)
//}
//
//func (self *Lockfile) LockWriteRange(offset int64, to int, len int64) error {
//	return self.lock(true, false, offset, to, len)
//}
//
//func (self *Lockfile) WaitAndLockReadRange(offset int64, to int, len int64) error {
//	return self.lock(false, true, offset, to, len)
//}
//
//func (self *Lockfile) WaitAndLockWriteRange(offset int64, to int, len int64) error {
//	return self.lock(true, true, offset, to, len)
//}
//
//func (self *Lockfile) UnlockRange(offset int64, to int, len int64) {
//	self.unlock(offset, to, len)
//}
//
//// Owner will return the pid of the process that owns an fcntl based
//// lock on the file. If the file is not locked it will return -1.
//func (self *Lockfile) Owner() int {
//	ft := &syscall.Flock_t{}
//	*ft = *self.ft
//
//	if err := syscall.FcntlFlock(self.file.Fd(), syscall.F_GETLK, ft); err != nil {
//		fmt.Println(err)
//		return -1
//		if ft.Type == syscall.F_UNLCK {
//			fmt.Println(err)
//			return -1
//		}
//	}
//	return int(ft.Pid)
//}
//
//func (self *Lockfile) lock(exclusive, blocking bool, offset int64, to int, len int64) error {
//	if self.file == nil {
//		f, err := os.OpenFile(self.Path, os.O_CREATE|os.O_RDWR, 0666)
//		if err != nil {
//			return err
//		}
//		self.file = f
//	}
//
//	ft := &syscall.Flock_t{
//		Whence: int16(to),
//		Start:  offset,
//		Len:    len,
//		Pid:    int32(os.Getpid()),
//	}
//	self.ft = ft
//
//	if exclusive {
//		ft.Type = syscall.F_WRLCK
//	} else {
//		ft.Type = syscall.F_RDLCK
//	}
//	var flags int
//	if blocking {
//		flags = syscall.F_SETLKW
//	} else {
//		flags = syscall.F_SETLK
//	}
//
//	if err := syscall.FcntlFlock(self.file.Fd(), flags, self.ft); err != nil {
//		if self.maintainFile {
//			self.file.Close()
//		}
//		return fmt.Errorf("failed to obtain lock")
//	}
//	return nil
//}
//
//func (self *Lockfile) unlock(offset int64, to int, len int64) {
//	self.ft.Len = len
//	self.ft.Start = offset
//	self.ft.Whence = int16(to)
//	self.ft.Type = syscall.F_UNLCK
//	syscall.FcntlFlock(self.file.Fd(), syscall.F_SETLK, self.ft)
//	if self.maintainFile {
//		self.file.Close()
//	}
//}
//
//type RangeLocker interface {
//	Locker
//	LockReadRange(int64, int, int64) error
//	LockWriteRange(int64, int, int64) error
//	WaitAndLockReadRange(int64, int, int64) error
//	WaitAndLockWriteRange(int64, int, int64) error
//	UnlockRange(int64, int, int64)
//}
//
//func NewLockfile(path string) RangeLocker {
//	return NewLockfile(path)
//}
//
//func NewLockfileFromFile(file *os.File) RangeLocker {
//	return NewLockfileFromFile(file)
//}
