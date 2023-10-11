package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func DropPrivs() error {
	uid, gid := GetUserEnvInt()
	if err := syscall.Setregid(-1, gid); err != nil {
		return errors.New(fmt.Sprintf("Can't drop gid: %s", err))
	}
	if err := syscall.Setreuid(-1, uid); err != nil {
		return errors.New(fmt.Sprintf("Can't drop uid: %s", err))
	}
	return nil
}

func EscalatePrivs() error {
	err := syscall.Setreuid(-1, 0)
	return err
}

// GetUserEnv checks if the process can drop priviledges by checking if either
// SUDO_UID and SUDO_GID or PKEXEC_UID are set, it returns the corresponding
// uid and gid set in these or 0 otherwise.
func GetUserEnv() (uid, gid string) {
	if v := os.Getenv("SUDO_UID"); v != "" {
		uid = v
	} else if v := os.Getenv("PKEXEC_UID"); v != "" {
		uid = v
	} else {
		uid = "0"
	}

	if v := os.Getenv("SUDO_GID"); v != "" {
		gid = v
	} else {
		gid = "0"
	}
	return uid, gid
}

func GetUserEnvInt() (uid, gid int) {
	uidString, gidString := GetUserEnv()
	uid, _ = strconv.Atoi(uidString)
	gid, _ = strconv.Atoi(gidString)
	return uid, gid
}
