package simple_test

import (
	"data/challenge-1-map/simple-slice-solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplekeyValue(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		t.Run("return true if key is present", func(t *testing.T) {
			sliceKeyValues := sliceKeyValuesWithLength(10)

			answer := sliceKeyValues.Has(5)
			assert.Equal(t, true, answer)
		})
		t.Run("return false if key is not present", func(t *testing.T) {
			sliceKeyValues := sliceKeyValuesWithLength(10)

			answer := sliceKeyValues.Has(100)
			assert.Equal(t, false, answer)
		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("gets the value if key present", func(t *testing.T) {
			expectedValue := "d"
			sliceKeyValues := sliceKeyValuesWithLength(10)

			value := sliceKeyValues.Get(3)

			assert.Equal(t, expectedValue, value)

		})
		t.Run("returns empty string if key not present", func(t *testing.T) {
			expectedValue := ""
			sliceKeyValues := sliceKeyValuesWithLength(10)

			value := sliceKeyValues.Get(100)

			assert.Equal(t, expectedValue, value)
		})
	})
	t.Run("Remove", func(t *testing.T) {
		t.Run("removes the element with given key and returns it", func(t *testing.T) {
			expectedRemovedElement := simple.SimpleKeyValue{
				Key:   3,
				Value: "d",
			}
			sliceKeyValues := sliceKeyValuesWithLength(10)

			removedElement := sliceKeyValues.Remove(3)
			assert.Equal(t, expectedRemovedElement, removedElement)

			assert.Len(t, sliceKeyValues.KeyValues, 9)
		})
		t.Run("if element is not present, returns nil", func(t *testing.T) {
			expectedRemovedElement := simple.SimpleKeyValue{
				Key:   0,
				Value: "",
			}
			sliceKeyValues := sliceKeyValuesWithLength(10)

			removedElement := sliceKeyValues.Remove(30)
			assert.Equal(t, expectedRemovedElement, removedElement)

			assert.Len(t, sliceKeyValues.KeyValues, 10)
		})
	})
	t.Run("index postition", func(t *testing.T) {
		sliceKeyValues := sliceKeyValuesWithLength(10)

		indexPosition := sliceKeyValues.IndexOf(3)
		assert.Equal(t, 3, indexPosition)
	})
}
func sliceKeyValuesWithLength(size int) *simple.SliceKeyValues {
	sliceKeyValues := simple.NewSliceKeyValues()
	initialSlice := newSlice(size)
	sliceKeyValues.KeyValues = initialSlice

	return sliceKeyValues
}
func newSlice(size int) []simple.SimpleKeyValue {
	var finalSlice []simple.SimpleKeyValue
	sliceSize := make([]simple.SimpleKeyValue, size)
	initialKey := 0
	initialValue := "z"
	for _, s := range sliceSize {
		s.Key = initialKey
		initialKey++
		s.Value = getNextString(initialValue)
		initialValue = s.Value

		finalSlice = append(finalSlice, s)
	}

	return finalSlice
}
func getNextString(str string) string {
	result := ""
	for _, letter := range str {
		if letter == 'z' {
			result += "a"
		} else {
			result += string(letter + 1)
		}
	}
	return result
}
