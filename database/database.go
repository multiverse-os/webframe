package database

// NOTE: Aim for this API
// Database("name").Put(...)
// TODO: I agree with this and not the mapb ullshit

import (
	keyvalue "github.com/multiverse-os/webframe/database/keyvalue"
)

// Alias
const (
	KV = KeyValue
)

const (
	Bitcask = keyvalue.Bitcask
	Pogreb  = keyvalue.Pogreb
	//BadgerDB = keyvalue.BadgerDB
)

// //////////////////////////////////////////////////////////////////////////////
type DatabaseType uint8

const (
	UndefinedDatabase DatabaseType = iota
	KeyValue
	//AppendOnly / TimeSeries
	// Something but forgot
	//Document
	//Columnar
	//Graph
	//ColumnOriented
	//RowOriented
)

// TODO: I dont want this here really, if I can MOVE this to framework and
// consolidate all the logic into there so we are calling in a single library
// from the application portion of the code

// TODO: Maybe we should be using nesting interfaces for Database abstracting
// the various TYPES of databases (KV, Column, Document, Graph)
type Database struct {
	Name      string
	Path      string
	Type      DatabaseType
	StoreType StoreType // derive path from this
	Store     keyvalue.Database
}

func OpenKV(kvType keyvalue.Type, storeType StoreType, path string) *Database {
	// TODO: Properly handle the error which may mean to make this not a
	// chain-able function (maybe make it possible for it not to error, by
	// providing defaults on faulty or blank entries)
	kv, _ := keyvalue.Open(kvType, storeType.String(), path)
	return &Database{
		Name:      storeType.String(),
		Path:      path,
		Type:      KeyValue,
		StoreType: storeType,
		Store:     kv,
	}
}

func OpenBitcask(name string) *Database {
	storeType := MarshalStoreType(name)
	return OpenKV(keyvalue.Bitcask, storeType, ("db/" + name))
}

func OpenPogreb(name string) *Database {
	storeType := MarshalStoreType(name)
	return OpenKV(keyvalue.Pogreb, storeType, ("db/" + name))
}
