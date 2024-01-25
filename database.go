package maglev

import (
	database "github.com/multiverse-os/maglev/database"
)

const (
	Data    = database.Data
	Session = database.Session
	Model   = database.Model
	Cache   = database.Cache
	Job     = database.Job
)

//func MarshalStoreType(typeStr string) database.StoreType {
//	switch typeStr {
//	case Session.String():
//		return Session
//	}
//}

// TODO: Need multiple k/v stores, one will be for session, one for cache, then
// maybe build models over k/v or just use an object db

//type Database struct {
//	*database.Database
//
//	Framework
//
//	Name string
//	Path string
//
//	Type database.StoreType
//}

func (f *Framework) KV(storeType database.StoreType) *database.Database {
	if f.databases[storeType] == nil {
		f.databases[storeType] = database.OpenBitcask(storeType.String())
	}
	return f.databases[storeType]
}

// TODO: Init to default db type for Model type
func (f *Framework) DB() *database.Database {
	if f.databases[database.Model] == nil {
		f.databases[database.Model] = database.OpenBitcask(database.Model.String())
	}
	return f.databases[database.Model]
}
