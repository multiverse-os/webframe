package webframe

import (
	database "github.com/multiverse-os/webframe/database"
	"github.com/multiverse-os/webframe/database/keyvalue"
)

const (
	UndefinedStoreType = database.UndefinedStoreType
	SessionStore       = database.Session
	ModelStore         = database.Model
	CacheStore         = database.Cache
)

// TODO: We have all 3 necessary for a basic user application app, while we
// still would like to be able to call them individually, a single call would be
// nice

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

// TODO: This database.Database sucks and why do we have a separate one here? It
// looks like this is all very undecided, it was switched up, and I added two KV
// databases.
func (f *Framework) KV(storeType database.StoreType) *database.Database {
	if f.databases[storeType] == nil {
		f.databases[storeType] = database.OpenBitcask(storeType.String())
	}
	return f.databases[storeType]
}

// TODO: This simplifies initializing each database.
func (f *Framework) InitializeDBs() *Framework {
	f.KV(ModelStore)
	f.KV(CacheStore)
	f.KV(SessionStore)
	return f
}

// TODO: Init to default db type for Model type
func (f *Framework) DB(storeType database.StoreType) *database.Database {
	if f.databases[storeType] == nil {
		f.databases[storeType] = database.OpenBitcask(storeType.String())
	}
	return f.databases[storeType]
}

// TODO: Id like to change the names, like app.Cache or app.Sessions
// and this should be maybe Database()
func (f *Framework) ModelDB() keyvalue.Database {
	return f.DB(ModelStore).Store
}

func (f *Framework) Cache() keyvalue.Database {
	return f.DB(CacheStore).Store
}

func (f *Framework) Sessions() keyvalue.Database {
	return f.DB(SessionStore).Store
}

//func (f *Framework) JobDB() *database.Database     {
// return f.DB(JobStore).Store
// }

//func (f *Framework) DataDB() *database.Database    {
// return f.DB(DataStore).Store
//}
