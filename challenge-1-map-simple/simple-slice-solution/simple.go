package simple

import (
	interfaces "data"
	"fmt"
)

type simpleKeyValue struct {
	key   int
	value string
}

type sliceKeyValues struct {
	keyValues []simpleKeyValue
}

func NewSliceKeyValues() interfaces.Operations {
	return &sliceKeyValues{}
}

func (s *sliceKeyValues) Has(key int) bool {
	if s.indexOf(key) == -1 {
		fmt.Printf("element with key: %v not present\n", key)
		return false
	}
	return true
}

func (s *sliceKeyValues) Get(key int) (string, bool) {
	for _, k := range s.keyValues {
		if k.key == key {
			return k.value, true
		}
	}
	fmt.Printf("element with key: %v not present\n", key)
	return "", false
}

func (s *sliceKeyValues) Remove(key int) (string, bool) {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		fmt.Printf("Key: %v, not found\n", key)
		return "", false
	}
	value, isPresent := s.Get(key)
	s.keyValues = append(s.keyValues[:keyIndex], s.keyValues[keyIndex+1:]...)

	fmt.Printf("removed element with key: %v and value: :%v\n", key, value)
	return value, isPresent
}

func (s *sliceKeyValues) Set(key int, value string) (string, bool) {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		s.keyValues = append(s.keyValues, simpleKeyValue{
			key:   key,
			value: value,
		})
		return "", true
	}
	oldValue, _ := s.Get(key)
	s.keyValues[keyIndex].value = value
	return oldValue, true
}

func (s *sliceKeyValues) indexOf(key int) int {
	for i, k := range s.keyValues {
		if k.key == key {
			return i
		}
	}
	return -1
}
