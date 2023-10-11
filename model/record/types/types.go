package types

type Type uint8

const (
	String Type = iota
	Id
	Rune
	Integer    // Unsigned, Size (uint64...) is stored in type struct
	Byte       // Single Byte
	Bytes      // Store images as base64 and use the data naturally
	Hash       // Hash, HashMap, Map, or Dictionary
	Enumerated // Functionality to work as a state machine included
	Decimal
	Float
	Boolean
)

// Aliasing DataType
const (
	Key          = Id
	Dictionary   = Hash
	Map          = Hash
	Money        = Decimal
	ByteSlice    = Bytes
	Data         = Bytes
	RuneSlice    = String
	Character    = Rune // UTF8 Character
	StateMachine = Enumerated
)
