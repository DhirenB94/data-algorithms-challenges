package simple_test

import (
	"data/challenge-1-map/simple-slice-solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplekeyValue(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		t.Run("return true if key is present", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()
			initialSlice := newSlice(10)
			sliceKeyValues.KeyValues = initialSlice

			answer := sliceKeyValues.Has(5)
			assert.Equal(t, true, answer)
		})
		t.Run("return false if key is not present", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()
			initialSlice := newSlice(10)
			sliceKeyValues.KeyValues = initialSlice

			answer := sliceKeyValues.Has(100)
			assert.Equal(t, false, answer)
		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("gets the value if key present", func(t *testing.T) {
			expectedValue := "d"
			sliceKeyValues := simple.NewSliceKeyValues()
			initialSlice := newSlice(10)
			sliceKeyValues.KeyValues = initialSlice

			value := sliceKeyValues.Get(3)

			assert.Equal(t, expectedValue, value)

		})
		t.Run("returns empty string if key not present", func(t *testing.T) {
			expectedValue := ""
			sliceKeyValues := simple.NewSliceKeyValues()
			initialSlice := newSlice(10)
			sliceKeyValues.KeyValues = initialSlice

			value := sliceKeyValues.Get(100)

			assert.Equal(t, expectedValue, value)
		})
	})
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

//extract to function
//newSlice := simple.NewSliceKeyValues()
//
//var oneTwo []simple.SimpleKeyValue
//one := simple.SimpleKeyValue{
//Key:   1,
//Value: "one",
//}
//two := simple.SimpleKeyValue{
//Key:   2,
//Value: "two",
//}
//
//oneTwo = append(oneTwo, one, two)
//newSlice.KeyValues = oneTwo
//
//fmt.Println(newSlice)
