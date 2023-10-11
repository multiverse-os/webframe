package record

type Id struct {
	Bytes []byte
}

func (i Id) String() string {
	return string(i.Bytes)
}
