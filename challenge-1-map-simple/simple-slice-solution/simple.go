package simple

import (
	interfaces "data"
	"fmt"
)

type simpleKeyValue struct {
	key   interface{}
	value interface{}
}

type sliceKeyValues struct {
	keyValues []simpleKeyValue
}

func NewSliceKeyValues() interfaces.Operations {
	return &sliceKeyValues{}
}

func (s *sliceKeyValues) Has(key interface{}) bool {
	if s.indexOf(key) == -1 {
		return false
	}
	return true
}

func (s *sliceKeyValues) Get(key interface{}) (interface{}, bool) {
	for _, k := range s.keyValues {
		if k.key == key {
			return k.value, true
		}
	}
	return nil, false
}

func (s *sliceKeyValues) Remove(key interface{}) (interface{}, bool) {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		return nil, false
	}
	value, isPresent := s.Get(key)
	s.keyValues = append(s.keyValues[:keyIndex], s.keyValues[keyIndex+1:]...)

	return value, isPresent
}

func (s *sliceKeyValues) Set(key interface{}, value interface{}) (interface{}, bool) {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		s.keyValues = append(s.keyValues, simpleKeyValue{
			key:   key,
			value: value,
		})
		return nil, true
	}
	oldValue, _ := s.Get(key)
	s.keyValues[keyIndex].value = value
	return oldValue, true
}

func (s *sliceKeyValues) indexOf(key interface{}) int {
	for i, k := range s.keyValues {
		if k.key == key {
			return i
		}
	}
	return -1
}

func (s *sliceKeyValues) String() string {
	return fmt.Sprint("KeyValues", s.keyValues)
}
