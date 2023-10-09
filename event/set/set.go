package set

import (
	"fmt"
	"sync"
)

// Set objects are collections of strings. A value in the Set may only occur once,
// it is unique in the Set's collection.
type Set struct {
	set   map[string]struct{}
	mutex *sync.Mutex
}

func (s *Set) Add(v string) error {
	s.init()
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.set == nil {
		s.set = make(map[string]struct{})
	}

	if _, ok := s.set[v]; ok {
		return fmt.Errorf("Value already in set.")
	}

	s.set[v] = struct{}{}
	return nil
}

func (s *Set) Clear() {
	s.init()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.set = nil
}

func (s *Set) Values() (keys []string) {
	s.init()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	keys = make([]string, 0, len(s.set))
	for k := range s.set {
		keys = append(keys, k)
	}

	return
}

func (s *Set) Has(v string) (ok bool) {
	s.init()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok = s.set[v]

	return
}

func (s *Set) init() {
	if s.mutex == nil {
		s.mutex = &sync.Mutex{}
	}
}
