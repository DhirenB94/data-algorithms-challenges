package simple

import (
	interfaces "data"
	"fmt"
)

type simpleKeyValue[K, V interfaces.CustomType] struct {
	key   K
	value V
}

type sliceKeyValues[K, V interfaces.CustomType] struct {
	keyValues []simpleKeyValue[K, V]
}

func NewSliceKeyValues[K, V interfaces.CustomType]() interfaces.Operations[K, V] {
	return &sliceKeyValues[K, V]{}
}

func (s *sliceKeyValues[K, V]) Has(key K) bool {
	if s.indexOf(key) == -1 {
		return false
	}
	return true
}

func (s *sliceKeyValues[K, V]) Get(key K) *V {
	for _, k := range s.keyValues {
		if k.key == key {
			return &k.value
		}
	}
	return nil
}

func (s *sliceKeyValues[K, V]) Remove(key K) *V {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		return nil
	}
	value := s.Get(key)
	s.keyValues = append(s.keyValues[:keyIndex], s.keyValues[keyIndex+1:]...)

	return value
}

func (s *sliceKeyValues[K, V]) Set(key K, value V) *V {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		s.keyValues = append(s.keyValues, simpleKeyValue[K, V]{
			key:   key,
			value: value,
		})
		return nil
	}
	oldValue := s.Get(key)
	s.keyValues[keyIndex].value = value
	return oldValue
}

func (s *sliceKeyValues[K, V]) indexOf(key K) int {
	for i, k := range s.keyValues {
		if k.key == key {
			return i
		}
	}
	return -1
}

func (s *sliceKeyValues[K, V]) String() string {
	return fmt.Sprint("KeyValues", s.keyValues)
}
