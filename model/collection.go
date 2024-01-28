package model

import (
	"bytes"
	"encoding/binary"
	"time"

	muid "github.com/multiverse-os/muid"
	database "github.com/multiverse-os/webframe/database"
)

//type SortDirection uint8
//
//const (
//  Unsorted SortDirection = iota
//	Ascending
//	Descending
//)
//
//type Sort struct {
//	Column    string
//	Direction SortDirection
//}
//
//type Index struct {
//	Name     string
//	Sorting  []*Sort
//	IsSorted bool
//	Records  []*Record
//}

type Collections []*Collection

// TODO: This is a good direction; just feel it out, it falls into place as you
// design it good.
type Collection struct {
	Model *Model

	Instance

	Name    string // TODO: Its currently 1 big key/value, but we will use the byte version of the name and prefix it to records keys with '#{collection.Name-#{id}'
	keys    []Key
	Records []*Record
	// Maybe like LastCreated, Last Modified; but these would provide us the
	// structure we need for tracking them and handles for various activites.

	// Also may solve GLOBAL problem; we have a database variable ( then each
	// model can be stored any way desirable (including using remote API
	// eventually).
	Database      *database.Database
	LastUpdatedAt time.Time
}

/// TODO: What if we could do New for collection ON THE BLOODY DATABASE OBJECT!
//Or the app object

// TODO: Yeah something like this needs to be called
// TODO: This is fucked,
func (db Database) NewCollection(name string) Collection {
	collection := Collection{
		Database: db,
		Name:     name,
	}

	//Collections = append(collections, &collection)
	return collection
}

func (c *Collection) UseDB(db *database.Database) Collection {
	c.Database = db
	return c
}

func (c Collection) Id() []byte {
	total := len(c.Name)
	for index, nameRune := range c.Name {
		total += index * int(nameRune)
	}
	binaryBuffer := new(bytes.Buffer)
	err := binary.Write(binaryBuffer, binary.LittleEndian, total)
	if err != nil {
		panic(err)
	}
	return binaryBuffer.Bytes()
}

// TODO: Would REALLy like a better way then appending full name, would prefer
// some sort of way to reduce to fixed length. or even an int we track to some
// like map. So we have a map in this package like Models when we create a new
// collection, then we append to Models, and that stores the Id we use to
// prefix. Ideally something like 0..99 but it would have to remain fixed across
// IDEA could be for example int value = len(self.Name) then ascii value for
// each position. This would give us a unique value that does not have overlap
// and it would be regeneratable. Length may seem unecesary but if you have
// different letter combos that add up the the same amount, using the length
// helps reduce the overlap; though doesnt resolve user = resu... so need
// something about position too.. len + (pos + ascii#)
//
//	user = 4 + (1 * 117) + (2 * 115) + (3 * 101) + (4 * 114) = 1110
//	resu = 4 + (1 * 114) + (2 * 101) + (3 * 115) + (4 * 117) = 686
func (c Collection) GenerateId() muid.Id {
	return GenerateId().Prefix(self.Name + ".")
}

//func (c Collection) Keys() {
//	for key := range Database.Store.Keys() {
//		if len(c.Name) < len(key) && bytes.Compare(key[:len(c.Name)], []byte(c.Name)) == 0 {
//			fmt.Printf("existing key IN collection: %v\n", string(key))
//		} else {
//			fmt.Printf("existing key NOT IN collection: %v\n", string(key))
//		}
//	}
//	fmt.Printf("got all keys!\n")
//
//}

//type Record struct {
//	Collection *Collection
//	Name       string
//	Id         record.Id
//	Data       record.Data
//}

// TODO: If we track changes then we can roll back, and depending on how much
//       we save we can selectively rollback

//// Collections ////////////////////////////////////////////////////////////////
//func (self *Collection) StepBackwards() (success bool, err error) {
//	// TODO: Take last item from the history, remove(pop) it. Then use this item
//	//       to revert the last change.
//
//	// TODO: Record should also have all its own updates stored within it, in
//	//       order so that it can be used to revert changes or show historical
//	//       versions using BSON diffing.
//  return false, nil
//}
//
//func (self *Collection) All() (records []*Record, err error) {
//	return self.Records, nil
//}
//
//func (self Collection) New(name string) (bool, *Collection) {
//	return false, nil
//}
//
//// Records ////////////////////////////////////////////////////////////////////
//func (self Collection) Delete(id record.Id) bool {
//	return false
//}
//
//func (self Collection) Find(id record.Id) (bool, *Record) {
//	return false, nil
//}
//
//func (self Collection) Update(id record.Id, record *Record) (bool, *Record) {
//	return false, nil
//}
//
//func (self Collection) Create(record *Record) (bool, *Record) {
//	return false, nil
//}
