package record

type Data struct {
	Bytes []byte
}

func (d Data) Key() string {
	// TODO: Extract key
	return ""
}

func (d Data) Value() []byte {
	// TODO Extract data
	return []byte{}
}

func (d Data) Object() interface{} {
	return nil
}

// TODO: Ability to parse out XML, JSON, BSON, etc
