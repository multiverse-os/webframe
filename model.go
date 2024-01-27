package webkit

import database "github.com/multiverse-os/webkit/database"

const (
	Uncompressed CompressionType = iota
	LZ
	ZStd
	Snappy
)

// TODO: Probably best to just find a binary to struct then modify it to support
//
//	our relevant types
const (
	Unencoded EncodingType = iota
	JSON
	CBOR
	BSON
)

// TODO: Then maybe we save serialize into BSON or CBOR to create our own object
//
//	store
type Model struct {
	Framework *Framework

	Database *database.Database

	Name string

	Key   []byte
	Value []byte

	Compression CompressionType
	Encoding    EncodingType

	BeforeCreate []func()
	AfterCreate  []func()

	BeforeSave []func()
	AfterSave  []func()
}
