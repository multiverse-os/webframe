package database

type Permissions uint8

const (
	WriteOnly Permissions = iota
	ReadOnly
	ReadAndWrite
	Disabled
)
