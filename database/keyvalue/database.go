package keyvalue

import (
	"fmt"
	//"errors"

	bitcask "git.mills.io/prologic/bitcask"
	pogreb "github.com/akrylysov/pogreb"
)

type Type int

const (
	Undefined Type = iota
	BadgerDB
	Pogreb
	Bitcask
	LevelDB
)

type Database interface {
	//Name() string
	//Path() string
	//Id() []byte

	// Bitcask
	//Stats() (bitcask.Stats, error)
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte) error
	Delete(key []byte) error

	//Len() int
	Close() error

	//Keys() chan []byte
	//Has(key []byte) bool
	//Sync() error
	//Backup(path string) error

	//PutWithTTL([]byte, []byte, time.Duration) error
	//Delete([]byte) error
	//Sift(func([]byte) (bool, error)) error
	//Scan([]byte, func([]byte) error) error
	//SiftScan([]byte, func([]byte) (bool, error)) error
	//Range([]byte, []byte, func([]byte) error) error
	//SiftRange([]byte, []byte, func([]byte) (bool, error)) error
	//RunGC() error
	//Fold(f func(key []byte) error) error
	//Reopen() error
	//Merge() error
	//Reclaimable() int64
}

func Open(databaseType Type, name, path string) (Database, error) {
	switch databaseType {
	case Bitcask:
		fmt.Printf("using bitcask name (not yet used): %v\n", name)
		return bitcask.Open(path)
	case Pogreb:
		fmt.Println("starting support for pogreb which is good for reads")
		return pogreb.Open(path, nil)
	case BadgerDB:
		fmt.Println("badger db not supported yet :(")
		return nil, nil
	default: // Undefined
		fmt.Println("undefined key value database")
		return nil, nil
	}
}
