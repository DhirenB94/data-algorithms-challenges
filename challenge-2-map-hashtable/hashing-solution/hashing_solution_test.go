package hashing_solution_test

import (
	"data/challenge-2-map-hashtable/hashing-solution"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHashKeys(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		t.Run("return true if key is present", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap()
			hashTable.Set(45, "foo")
			hashTable.Set(95, "bar")

			timeNow := time.Now()
			has := hashTable.Has(45)
			fmt.Println(time.Since(timeNow))

			assert.True(t, has)

			has = hashTable.Has(95)
			assert.True(t, has)
		})
		t.Run("return false if key is not present", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap()

			has := hashTable.Has(10)
			assert.False(t, has)
		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("returns the value and true if key present", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap()
			expectedGetValue := "foo"
			expectedGetValue2 := "bar"
			hashTable.Set(45, "foo")
			hashTable.Set(95, "bar")

			timeNow := time.Now()
			getValue, isGot := hashTable.Get(45)
			fmt.Println(time.Since(timeNow))

			assert.Equal(t, expectedGetValue, getValue)
			assert.True(t, isGot)

			getValue, isGot = hashTable.Get(95)
			assert.Equal(t, expectedGetValue2, getValue)
			assert.True(t, isGot)
		})
		t.Run("returns empty string and false if key not present", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap()

			getValue, isGot := hashTable.Get(45)
			assert.Equal(t, "", getValue)
			assert.False(t, isGot)
		})
	})
	t.Run("Set", func(t *testing.T) {
		t.Run("returns empty string and true if element does not exist already", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap()

			timeNow := time.Now()
			oldValue, isSet := hashTable.Set(45, "foo")
			fmt.Println(time.Since(timeNow))

			assert.Empty(t, oldValue)
			assert.True(t, isSet)

			getValue, isGot := hashTable.Get(45)
			assert.Equal(t, "foo", getValue)
			assert.True(t, isGot)

			oldValue, isSet = hashTable.Set(95, "bar")
			assert.Empty(t, oldValue)
			assert.True(t, isSet)

			getValue, isGot = hashTable.Get(95)
			assert.Equal(t, "bar", getValue)
			assert.True(t, isGot)
		})
		t.Run("if element with key exists already, update its value, return the old value and true", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap()

			timeNow := time.Now()
			oldValue, isSet := hashTable.Set(45, "foo")
			fmt.Println(time.Since(timeNow))

			assert.Empty(t, oldValue)
			assert.True(t, isSet)

			oldValue, isSet = hashTable.Set(45, "bar")
			assert.Equal(t, "foo", oldValue)

			getValue, _ := hashTable.Get(45)
			assert.Equal(t, "bar", getValue)
		})
	})
	t.Run("Remove", func(t *testing.T) {
		t.Run("remove key if exists at the head of linked list", func(t *testing.T) {
			hashMap := hashing_solution.NewHashMap()

			hashMap.Set(45, "foo")
			hashMap.Set(95, "foo2")
			hashMap.Set(145, "foo3")

			removedValue, isRemoved := hashMap.Remove(45)
			assert.Equal(t, "foo", removedValue)
			assert.True(t, isRemoved)

			getValue, isGot := hashMap.Get(45)
			assert.Empty(t, getValue)
			assert.False(t, isGot)
		})
		t.Run("remove key if exists somewhere in the linked list", func(t *testing.T) {
			hashMap := hashing_solution.NewHashMap()

			hashMap.Set(45, "foo")
			hashMap.Set(95, "foo2")
			hashMap.Set(145, "foo3")
			hashMap.Set(195, "foo4")

			timeNow := time.Now()
			removedValue, isRemoved := hashMap.Remove(95)
			fmt.Println(time.Since(timeNow))

			assert.Equal(t, "foo2", removedValue)
			assert.True(t, isRemoved)

			getValue, isGot := hashMap.Get(55)
			assert.Empty(t, getValue)
			assert.False(t, isGot)
		})
		t.Run("return empty string and false if key does not exist", func(t *testing.T) {
			hashMap := hashing_solution.NewHashMap()

			hashMap.Set(45, "foo")
			hashMap.Set(10, "foo2")

			removedValue, isRemoved := hashMap.Remove(22)
			assert.Empty(t, removedValue)
			assert.False(t, isRemoved)
		})
	})
}
