package types

import (
	"fmt"
)

type Value struct {
	DataType Type
	Data     interface{}
}

func (self *Value) String() string {
	switch self.DataType {
	case String:
		return self.Data.(string)
	case Integer:
		// TODO: Convert integer into string
		// TODO: Use int data to determine which kind of int we are typecasting to
		return fmt.Sprintf("%v", self.Data.(int64))
	default:
		return ""
	}
}

// TODO: Integer, Hash/Dictionary/Map, Enumerated, Decimal, Money, etc
//       each should have the ability to convert some of the other
//       datatypes but should focus on ensuring the data is correct and this
//       will most likely serve as the place we transform data, conform data,
//       and ensure the user input is valid.
