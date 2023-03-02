package simple

import "fmt"

type SimpleKeyValue struct {
	Key   int
	Value string
}

type SliceKeyValues struct {
	KeyValues []SimpleKeyValue
}

func NewSliceKeyValues() *SliceKeyValues {
	return &SliceKeyValues{}
}

func (s *SliceKeyValues) Has(key int) bool {
	if s.IndexOf(key) == -1 {
		fmt.Printf("element with key: %v not present\n", key)
		return false
	}
	return true
}

func (s *SliceKeyValues) Get(key int) string {
	for _, k := range s.KeyValues {
		if k.Key == key {
			return k.Value
		}
	}
	fmt.Printf("element with key: %v not present\n", key)
	return ""
}
func (s *SliceKeyValues) Remove(key int) SimpleKeyValue {
	keyIndex := s.IndexOf(key)
	value := s.Get(key)
	if keyIndex == -1 {
		fmt.Printf("Key: %v, not found\n", key)
		return SimpleKeyValue{}
	}
	s.KeyValues = append(s.KeyValues[:keyIndex], s.KeyValues[keyIndex+1:]...)

	fmt.Printf("removed element with key: %v and value: :%v\n", key, value)
	return SimpleKeyValue{
		Key:   key,
		Value: value,
	}
}

func (s SliceKeyValues) IndexOf(key int) int {
	for i, k := range s.KeyValues {
		if k.Key == key {
			return i
		}
	}
	return -1
}
