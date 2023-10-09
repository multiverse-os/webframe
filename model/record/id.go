package record

type Id struct {
	Bytes []byte
}

func (self Id) String() string {
	return string(self.Bytes)
}
