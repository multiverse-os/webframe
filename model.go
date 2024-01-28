package webframe

import "github.com/multiverse-os/webframe/model"

type Model struct {
	*model.Collection

	Framework *Framework
}

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
