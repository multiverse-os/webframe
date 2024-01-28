package webframe

import (
	database "github.com/multiverse-os/webframe/database"
)

const (
	UndefinedStoreType = database.UndefinedStoreType
	SessionStore       = database.Session
	ModelStore         = database.Model
	CacheStore         = database.Cache
)

//func MarshalStoreType(typeStr string) database.StoreType {
//	switch typeStr {
//	case Session.String():
//		return Session
//	}
//}

// TODO: Need multiple k/v stores, one will be for session, one for cache, then
// maybe build models over k/v or just use an object db

type Database struct {
	Framework *Framework

	//*database.Database

	Name string
	Path string

	StoreType database.StoreType
}

func (f *Framework) KV(storeType database.StoreType) *database.Database {
	if f.databases[storeType] == nil {
		f.databases[storeType] = database.OpenBitcask(storeType.String())
	}
	return f.databases[storeType]
}

// TODO: Init to default db type for Model type
func (f *Framework) DB(storeType database.StoreType) *database.Database {
	if f.databases[storeType] == nil {
		f.databases[storeType] = database.OpenBitcask(storeType.String())
	}
	return f.databases[storeType]
}

func (f *Framework) ModelDB() *database.Database   { return f.DB(ModelStore) }
func (f *Framework) CacheDB() *database.Database   { return f.DB(CacheStore) }
func (f *Framework) SessionDB() *database.Database { return f.DB(SessionStore) }

//func (f *Framework) JobDB() *database.Database     { return f.DB(JobStore) }
//func (f *Framework) DataDB() *database.Database    { return f.DB(DataStore) }
