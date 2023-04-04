package simple

import (
	interfaces "data"
	"fmt"
)

type simpleKeyValue[V any] struct {
	key   int
	value V
}

type sliceKeyValues[V any] struct {
	keyValues []simpleKeyValue[V]
}

func NewSliceKeyValues[V any]() interfaces.Operations[V] {
	return &sliceKeyValues[V]{}
}

func (s *sliceKeyValues[V]) Has(key int) bool {
	if s.indexOf(key) == -1 {
		return false
	}
	return true
}

func (s *sliceKeyValues[V]) Get(key int) *V {
	for _, k := range s.keyValues {
		if k.key == key {
			return &k.value
		}
	}
	return nil
}

func (s *sliceKeyValues[V]) Remove(key int) *V {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		return nil
	}
	value := s.Get(key)
	s.keyValues = append(s.keyValues[:keyIndex], s.keyValues[keyIndex+1:]...)

	return value
}

func (s *sliceKeyValues[V]) Set(key int, value V) *V {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		s.keyValues = append(s.keyValues, simpleKeyValue[V]{
			key:   key,
			value: value,
		})
		return nil
	}
	oldValue := s.Get(key)
	s.keyValues[keyIndex].value = value
	return oldValue
}

func (s *sliceKeyValues[V]) indexOf(key int) int {
	for i, k := range s.keyValues {
		if k.key == key {
			return i
		}
	}
	return -1
}

func (s *sliceKeyValues[V]) String() string {
	return fmt.Sprint("KeyValues", s.keyValues)
}
