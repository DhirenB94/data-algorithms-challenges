package simple

import (
	"fmt"
)

type Operations interface {
	Has(key int) bool
	Get(key int) (string, bool)
	Remove(key int) (string, bool)
	Set(key int, value string) string
}

type SimpleKeyValue struct {
	Key   int
	Value string
}

type SliceKeyValues struct {
	KeyValues []SimpleKeyValue
}

func NewSliceKeyValues() Operations {
	return &SliceKeyValues{make([]SimpleKeyValue, 0)}
}

func (s *SliceKeyValues) Has(key int) bool {
	if s.indexOf(key) == -1 {
		fmt.Printf("element with key: %v not present\n", key)
		return false
	}
	return true
}

func (s *SliceKeyValues) Get(key int) (string, bool) {
	for _, k := range s.KeyValues {
		if k.Key == key {
			return k.Value, true
		}
	}
	fmt.Printf("element with key: %v not present\n", key)
	return "", false
}

func (s *SliceKeyValues) Remove(key int) (string, bool) {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		fmt.Printf("Key: %v, not found\n", key)
		return "", false
	}
	value, isPresent := s.Get(key)
	s.KeyValues = append(s.KeyValues[:keyIndex], s.KeyValues[keyIndex+1:]...)

	fmt.Printf("removed element with key: %v and value: :%v\n", key, value)
	return value, isPresent
}

func (s *SliceKeyValues) Set(key int, value string) string {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		s.KeyValues = append(s.KeyValues, SimpleKeyValue{
			Key:   key,
			Value: value,
		})
		return ""
	}
	oldValue, _ := s.Get(key)
	s.KeyValues[keyIndex].Value = value
	return oldValue
}

func (s *SliceKeyValues) indexOf(key int) int {
	for i, k := range s.KeyValues {
		if k.Key == key {
			return i
		}
	}
	return -1
}
