package simple_test

import (
	"data/challenge-1-map-simple/simple-slice-solution"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSimplekeyValue(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		t.Run("returns true if element is present", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()

			sliceKeyValues.Set(5, "foo")
			sliceKeyValues.Set(10, "bar")

			has := sliceKeyValues.Has(10)
			assert.True(t, has)
		})
		t.Run("returns false if element is not present", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()

			sliceKeyValues.Set(5, "foo")
			sliceKeyValues.Set(10, "bar")

			has := sliceKeyValues.Has(100)
			assert.False(t, has)
		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("gets the value if key present", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()
			expectedValue := "bar"

			sliceKeyValues.Set(45, "foo")
			sliceKeyValues.Set(50, "bar")

			timeNow := time.Now()
			value, isPresent := sliceKeyValues.Get(50)
			fmt.Println(time.Since(timeNow))

			assert.True(t, isPresent)
			assert.Equal(t, expectedValue, value)
		})
		t.Run("returns empty string if key not present", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()

			sliceKeyValues.Set(45, "foo")
			sliceKeyValues.Set(50, "bar")

			value, isPresent := sliceKeyValues.Get(100)
			assert.False(t, isPresent)
			assert.Empty(t, value)
		})
	})
	t.Run("Remove", func(t *testing.T) {
		t.Run("removes the element with given key", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()
			expectedRemovedValue := "foo"

			sliceKeyValues.Set(45, "foo")
			sliceKeyValues.Set(50, "bar")

			timeNow := time.Now()
			removedValue, isPresent := sliceKeyValues.Remove(45)
			fmt.Println(time.Since(timeNow))

			assert.True(t, isPresent)
			assert.Equal(t, expectedRemovedValue, removedValue)

			has := sliceKeyValues.Has(45)
			assert.False(t, has)
		})
		t.Run("if element is not present, returns nil", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()

			sliceKeyValues.Set(45, "foo")
			sliceKeyValues.Set(50, "bar")

			removedElement, isPresent := sliceKeyValues.Remove(100)

			assert.False(t, isPresent)
			assert.Empty(t, removedElement)
		})
	})
	t.Run("Set", func(t *testing.T) {
		t.Run("returns an empty string and true if element does not exist, and adds the element", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()

			has := sliceKeyValues.Has(100)
			assert.False(t, has)

			timeNow := time.Now()
			oldValue, isSet := sliceKeyValues.Set(100, "zzz")
			fmt.Println(time.Since(timeNow))

			assert.Empty(t, oldValue)
			assert.True(t, isSet)

			has = sliceKeyValues.Has(100)
			assert.True(t, has)
		})
		t.Run("if element with key exists already, update its value", func(t *testing.T) {
			sliceKeyValues := simple.NewSliceKeyValues()
			initialValue := "abc"
			expectedNewValue := "def"

			sliceKeyValues.Set(100, "abc")

			timeNow := time.Now()
			oldValue, isSet := sliceKeyValues.Set(100, "def")
			fmt.Println(time.Since(timeNow))

			assert.Equal(t, oldValue, initialValue)
			assert.True(t, isSet)

			actualNewValue, _ := sliceKeyValues.Get(100)

			assert.Equal(t, expectedNewValue, actualNewValue)
		})
	})
}
