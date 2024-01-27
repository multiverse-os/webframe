package database

// Basic enumerator
type StoreType int // should be uint8 but wait till later stage (bc stream gievs progression and chances to show how to imrpove things!!

const (
	UndefinedStoreType StoreType = iota
	Session
	Model
	Cache
	//Job
)

func MarshalStoreType(typeStr string) StoreType {
	switch typeStr {
	case Session.String():
		return Session
	case Model.String():
		return Model
	case Cache.String():
		return Cache
	//case Job.String():
	//	return Job
	default: // Undefined
		return UndefinedStoreType
	}
}

func (st StoreType) String() string {
	switch st {
	case Session:
		return "session"
	case Model:
		return "model"
	case Cache:
		return "cache"
	//case Job:
	//	return "job"
	default: // Undefined
		return "undefined"
	}
}
