package model

import "time"

//import (
//	"time"
//
//  record "github.com/multiverse-os/webkit/model/record"
//)
//
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

type Collection struct {
	Model *Model

	Name string
	//Records       []*Record
	//Indexes       []*Index
	LastUpdatedAt time.Time
	//History       chan *Record
}

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
