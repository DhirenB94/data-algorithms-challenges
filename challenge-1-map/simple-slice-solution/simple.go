package simple

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
	for _, k := range s.KeyValues {
		if k.Key == key {
			return true
		}
	}
	return false
}

func (s *SliceKeyValues) Get(key int) string {
	for _, k := range s.KeyValues {
		if k.Key == key {
			return k.Value
		}
	}
	return ""
}
func (s *SliceKeyValues) Remove(key int) SimpleKeyValue {
	for _, k := range s.KeyValues {
		if k.Key == key {
			return k
		}
	}
	return SimpleKeyValue{}
}
