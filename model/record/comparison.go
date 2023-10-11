package record

type Operator uint8
type Value interface{}

const (
	Is Operator = iota
	IsNot
	IsLess
	IsLessOrEqual
	IsMore
	IsMoreOrEqual
)

// Aliasing CompareType
const (
	IsEqual          = Is
	IsNotEqual       = IsNot
	IsGreater        = IsMore
	IsGreaterOrEqual = IsMoreOrEqual
)

type Comparison struct {
	Operator Operator
	Value    Value
}
