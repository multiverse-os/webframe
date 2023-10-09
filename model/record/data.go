package record

type Data struct {
	Bytes      []byte
}

func (self Data) Key() string {
  // TODO: Extract key
  return ""
}

func (self Data) Value() []byte {
  // TODO Extract data
  return []byte{}
}

func (self Data) Object() interface{} {
  return nil
}

// TODO: Ability to parse out XML, JSON, BSON, etc
