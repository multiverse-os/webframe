package model

type DataType uint

const (
	UndefinedType DataType = iota
	StringType
	FloatType
	BoolType
	RuneType
	ByteType
	ArrayType
	HashType
	IntegerType
)

type Attribute struct {
	Name    string
	Type    DataType
	Default string
	// TODO: Validation should definitely be built into the attribute itslef
	//       then we could have individual errors for each, but there will be
	//       some validations that require more than one attribute to determine
	// TODO: Would be nice to have constraints, like validations maybe, could
	//       do relations without a separate table, by just namingt he
	//       collection then collection:id; its one less hop, could be 1 to many
	//       collection:id,id,id;
	// Built into concept of tree or relationship structure would be nice
}
