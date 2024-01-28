package webframe

import "github.com/multiverse-os/webframe/model"

// TODO: Looks like much of this should be stuck in model folder so its
// model.Thing instead of framework.Model framework.Attribute etc

// TODO: Just look to CLI data, we already started on a simplified type system
// for situations like this, get us close enough then just handle the rest in
// application

// TODO: Then maybe we save serialize into BSON or CBOR to create our own object
//
//	store
//
// TODO: We will want the ability to flatten a model with embedded or
// relationships, so we can edit a single value and re-construct the object
// without needing to overwrite everything, if we keep everything bitwise
// separate we can do this easier.

// TODO: Honestly this logic really should go into model if we want a
// an independent which would be ideal, then this logic should go in database
// or model, and then model in database, to create basically an ORM that is
// very active record like, but its also okay if its not independent

type Model *model.Collection

//type Model struct {
//	Framework *Framework
//	Database  *database.Database
//
//	// TODO: This is our crude way of trying to work in a linked list
//	// functionality at the base level
//	Index uint
//	// TODO: We will encode the timestamp into the ID so we can sort by ID and
//	// get time sorting.
//	ID   []byte
//	Name string
//
//	// TODO: What else would be really common? state machine?
//
//	//  the obvious thing would seem to be methods; like User.Find(user_id) =>
//	//  User{id: user_id}
//	// If we store methods, it does make it easier to lock them down by saying
//	// which controllers can even use certain methods
//
//	Methods map[string]*Method
//
//	Attributes map[string]*Attribute
//
//	BeforeCreate []func()
//	AfterCreate  []func()
//
//	BeforeSave []func()
//	AfterSave  []func()
//}
//
//func (m *Model) Attribute(attributeName string, attribute *Attribute) *Model {
//	m.Attributes[attributeName] = attribute
//	return m
//}
//
//func (m *Model) Method(methodName string, method *Method) *Model {
//	m.Methods[methodName] = method
//	return m
//}
