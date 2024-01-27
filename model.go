package webkit

import database "github.com/multiverse-os/webkit/database"

// TODO: Just look to CLI data, we already started on a simplified type system
// for situations like this, get us close enough then just handle the rest in
// application
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

//type Validation struct {
//	Attribute *Attribute
//}

type Method func()

// TODO: Then maybe we save serialize into BSON or CBOR to create our own object
//
//	store
//
// TODO: We will want the ability to flatten a model with embedded or
// relationships, so we can edit a single value and re-construct the object
// without needing to overwrite everything, if we keep everything bitwise
// separate we can do this easier.
type Model struct {
	Framework *Framework
	Database  *database.Database

	// TODO: This is our crude way of trying to work in a linked list
	// functionality at the base level
	Index uint
	// TODO: We will encode the timestamp into the ID so we can sort by ID and
	// get time sorting.
	ID   []byte
	Name string

	// TODO: What else would be really common? state machine?

	//  the obvious thing would seem to be methods; like User.Find(user_id) =>
	//  User{id: user_id}
	// If we store methods, it does make it easier to lock them down by saying
	// which controllers can even use certain methods

	Methods map[string]*Method

	Attributes map[string]*Attribute

	BeforeCreate []func()
	AfterCreate  []func()

	BeforeSave []func()
	AfterSave  []func()
}

func (m *Model) Attribute(attributeName string, attribute *Attribute) *Model {
	m.Attributes[attributeName] = attribute
	return m
}

func (m *Model) Method(methodName string, method *Method) *Model {
	m.Methods[methodName] = method
	return m
}
