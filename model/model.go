package model

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"time"

	database "github.com/multiverse-os/maglev/database"

	// TODO: Should be submodule, ideally all packages will be submodules so there
	// are no third party dependencies.
	muid "github.com/multiverse-os/muid"
)

type Database struct {
	*database.Database

	Collections Collections
}

func GenerateId() muid.Id {
	return muid.Generate().Base32()
}

type Key []byte

// TODO: Should be able to get the collection object from any key but for now...
func (self Key) CollectionName() string {
	idBytes := []byte{}
	for _, idByte := range self {
		// "." == 46
		if int(idByte) == 46 {
			break
		}
		idBytes = append(idBytes, idByte)
	}
	return string(idBytes)
}

// TODO: Build methods for key, to filter, conversion, etc. (remember we are
// using muid.Id, so we should be able to get the bytes to muid id type so we
// can extract out the timestamp for example

// TODO: This is a good direction; just feel it out, it falls into place as you
// design it good.
type Collection struct {
	Instance

	Name    string // TODO: Its currently 1 big key/value, but we will use the byte version of the name and prefix it to records keys with '#{collection.Name-#{id}'
	keys    []Key
	Records []*Record
	// Maybe like LastCreated, Last Modified; but these would provide us the
	// structure we need for tracking them and handles for various activites.

	// Also may solve GLOBAL problem; we have a database variable ( then each
	// model can be stored any way desirable (including using remote API
	// eventually).
	Database *database.Database
}

type Collections []*Collection

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

func (c Collection) UseDB(db *database.Database) Collection {
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

type Instance interface{}

func Encode(obj interface{}) ([]byte, error) {
	byteBuffer := new(bytes.Buffer)
	if err := gob.NewEncoder(byteBuffer).Encode(obj); err != nil {
		return nil, err
	}
	return byteBuffer.Bytes(), nil
}

func Decode(data []byte, value interface{}) error {
	buf := bytes.NewBuffer(data)
	return gob.NewDecoder(buf).Decode(value)
}

//type Record struct {
//	Collection *Collection
//	Name       string
//	Id         record.Id
//	Data       record.Data
//}

type Record struct {
	Collection
	Instance

	Id        muid.Id
	CreatedAt time.Time
	UpdatedAt time.Time
	Binary    []byte // Store the data in binary format, which we convert in and out of for k/v
}

// TODO: So this Model system will provide for a way to convert binary data
// between our struct. So we need to have the correct codec to handle the gob.

// TODO: Okay so KEY CONCEPT here is we need to get saving working. So we create
// our generic model. And it already has through embedding a way to do .Save()
// on the object and it gets inserted into the database. This would be a huge
// step in creating a model concept and it would be far better than what I have
// seen just starting with that alone. WE NEED to finish the MVC model, basic
// MVC functionality would actually get us near alpha release and put us at our
// furtherest point after several refactors. it would be the skeleton we could
// shave off portions and clean up the rest to get it into a lean web framework;
// and we need to make it HTTP server neutral (gin, chi, echo, etc). So we will
// need to create an interface for that, and that will get us our rails like
// conrollers. Then we just build out what we are already doing in views to the
// poin that we ahve so many helpers that building out views is incredibly
// simple and basically no html ever.
//
//	self.Save(self) It would be better at that point to be model.Save(u1)
func (r *Record) GenerateId() muid.Id {
	if r.Id.IsNil() {
		r.Id = r.Collection.GenerateId()
	}
	return r.Id
}

func (r *Record) Save() bool {
	fmt.Printf("*Instance(r) %v\n", r)
	fmt.Printf("*Instance(r.Id) %v\n", r.Id)
	fmt.Printf("*Instance(r.Instance) %v\n", r.Instance)

	r.GenerateId()
	//fmt.Printf("*Instance(self.Name) %v\n", self.User.Name) - this should work,
	//I see the username in printing self but I cant access the field

	// TODO: Okay, so
	//  * We encode to GOB (for now) and cache our binary in the []byte
	//  * Then we create a bsonid like ID using muid (should happen @ new())
	//  * Then we use the muid in Bytes() form to issue the put.
	//      [[ but this doesnt give us our 'collection' or anything,
	//         maybe look at our insistance of putting in prefix
	//         despite wanting to throw it out 100 times ]]
	//  * Then we need to have the ability to pull these records from id
	//    and we can probably do searching by just loading up a collection
	//    and searching through it, because that would be loading from disk
	//    to memory to search, which is good.

	return true
}

func (r *Record) New() bool {
	r.CreatedAt = time.Now()
	r.UpdatedAt = r.CreatedAt
	r.GenerateId()

	// TODO: Actuaolly use thje GOB to convert it down; and hey why not while we
	// are at it do a compression. We can go back later and use like our codec
	// library hat is like totally fucking good and then wwe can streamline this
	// and use like cbor or something instead of gob.

	// But first we do that and store it in the byte slice named like data or raw.
	// Then we take that byte slice and save it in our k/v db under the id that
	// ahs our prefix so we can select out collections.

	recordBytes, _ := Encode(r)

	fmt.Printf("recordBytes(%v)", recordBytes)

	result := r.Collection.Database.Store.Put(r.Id.Bytes(), recordBytes)
	fmt.Printf("was there a result?(%v)", result)

	return true
}

/////////////////////////////////////////////////////////////////////////
// Model as a non-interface would mean we do inheritance through embedding or
// delegation (one way of saying it in golang)
// type User struct {
//   Model
// }

// One attempt below
//
//type Model struct {
//	Name           string
//	Collection     string
//	Fields         []Field
//}
//
//type Field struct {
//  DataType   string // TODO: Probably better enumerator we define for specialized types
//	Index      bool
//	PrimaryKey bool
//	References Model
//	Check      func(Model) bool
//	ColumnName string
//	Unique     bool
//}
//
//func (f *Field) Validate() bool {
//	return true
//}
//
//// returns a new field.  We should be able to take a hash of the attributes
//// and convert into the field directly
//func NewField() *Field {
//	return &Field{}
//}
//
//
//type DatabasePayload struct {
//}
//
//// generates a new model
//func NewModel(fields map[string]string) *Model {
//	model := &Model{}
//	return model
//}
//
//
/////
//
//type ModelInstance struct {
//	Model
//	Values    []Mapping
//	Columns   []string
//	RowNumber int
//	Errors    []errors.DatabaseOpError
//}
//
//type Mapping struct {
//	Field
//	Value interface{}
//}
//
//// func (m *ModelInstance) Save(options map[string]interface{}) bool {
//
//// }
//
//func (m *ModelInstance) Valid() bool {
//	if len(m.Values) == 0 {
//		return false
//	}
//	for _, mapping := range m.Values {
//		if typeOf(mapping.Value) != mapping.Field.DataType {
//			return false
//		}
//	}
//	return true
//}
//
//func (m *ModelInstance) ToJson() {
//
//}
//
//func (m *ModelInstance) BuildAttributes(attrs map[string]interface{}) {
//	mappings := []Mapping{}
//	for _, field := range m.Fields {
//		mapping := Mapping{
//			Field: field,
//			Value: attrs[field.ColumnName],
//		}
//		mappings = append(mappings, mapping)
//	}
//	m.Values = mappings
//}
//
//func typeOf(val interface{}) string {
//	return reflect.TypeOf(val).Kind().String()
//}
