package record

type DataType int

// // TODO: We may suffer by simplifying datatypes to the activerecord ones
const (
	String DataType = iota
	Integer
	Float
	Decimal
	Datetime
	Boolean
	Byte
	Map
	Slice
)
