package store

// Basic enumerator
type Type int // should be uint8 but wait till later stage (bc stream gievs progression and chances to show how to imrpove things!!

const (
  Data Type = iota
  Session 
  Model
  Cache
  Job
)

func MarshalType(typeStr string) Type {
  switch typeStr {
  case Session.String():
    return Session
  case Model.String():
    return Model
  case Cache.String():
    return Cache
  case Job.String():
    return Job
  default:
    return Data
  }
}

func (self Type) String() string {
  switch self {
  case Session:
    return "session"
  case Model:
    return "model"
  case Cache:
    return "cache"
  case Job:
    return "job"
  default: // Data
    return "data"
  }
}

