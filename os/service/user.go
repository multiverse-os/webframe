package service

import (
	"os"
	"os/user"
	"strconv"
)

type Group struct {
	Name string
	Gid  string
}

type User struct {
	Name     string
	Uid      string
	GroupIds []string
	Groups   []Group
	Home     string
	Root     bool
}

func CurrentUser() (*User, error) {
	if user, err := user.Current(); err != nil {
		return nil, err
	} else {
		groupIds, _ := user.GroupIds()
		return &User{
			Name:     user.Name,
			Uid:      user.Uid,
			GroupIds: groupIds,
			Root:     (user.Uid == strconv.Itoa(0)),
			Home:     os.Getenv("HOME"),
		}, nil
	}
}

func IsRootUser() bool {
	return (os.Geteuid() == 0)
}

func IsUnpriviledgedUser() bool {
	return !IsRootUser()
}
