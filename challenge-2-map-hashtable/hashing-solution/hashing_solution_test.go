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
			hashTable := hashing_solution.NewHashMap[string]()
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
			hashTable := hashing_solution.NewHashMap[string]()

			has := hashTable.Has(10)
			assert.False(t, has)
		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("returns the value if key present", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap[string]()
			expectedGetValue := "foo"
			expectedGetValue2 := "bar"
			hashTable.Set(45, "foo")
			hashTable.Set(95, "bar")

			timeNow := time.Now()
			getValue := hashTable.Get(45)
			fmt.Println(time.Since(timeNow))

			assert.Equal(t, &expectedGetValue, getValue)

			getValue = hashTable.Get(95)
			assert.Equal(t, &expectedGetValue2, getValue)
		})
		t.Run("returns nil if key not present", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap[string]()

			getValue := hashTable.Get(45)
			assert.Nil(t, getValue, getValue)
		})
	})
	t.Run("Set", func(t *testing.T) {
		t.Run("set the element and return nil if element does not exist already", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap[string]()

			timeNow := time.Now()
			oldValue := hashTable.Set(45, "foo")
			fmt.Println(time.Since(timeNow))

			assert.Nil(t, oldValue)

			getValue := hashTable.Get(45)
			assert.Equal(t, "foo", *getValue)

			oldValue = hashTable.Set(95, "bar")
			assert.Nil(t, oldValue)

			getValue = hashTable.Get(95)
			assert.Equal(t, "bar", *getValue)
		})
		t.Run("if element with key exists already, update its value but return the old value", func(t *testing.T) {
			hashTable := hashing_solution.NewHashMap[string]()

			timeNow := time.Now()
			oldValue := hashTable.Set(45, "foo")
			fmt.Println(time.Since(timeNow))

			assert.Nil(t, oldValue)

			oldValue = hashTable.Set(45, "bar")
			assert.Equal(t, "foo", *oldValue)

			getValue := hashTable.Get(45)
			assert.Equal(t, "bar", *getValue)
		})
	})
	t.Run("Remove", func(t *testing.T) {
		t.Run("remove key if exists at the head of linked list and return the removed value", func(t *testing.T) {
			hashMap := hashing_solution.NewHashMap[string]()

			hashMap.Set(45, "foo")
			hashMap.Set(95, "foo2")
			hashMap.Set(145, "foo3")

			removedValue := hashMap.Remove(45)
			assert.Equal(t, "foo", *removedValue)

			getValue := hashMap.Get(45)
			assert.Nil(t, getValue)
		})
		t.Run("remove key if exists somewhere in the linked list and return the removed value", func(t *testing.T) {
			hashMap := hashing_solution.NewHashMap[string]()

			hashMap.Set(45, "foo")
			hashMap.Set(95, "foo2")
			hashMap.Set(145, "foo3")
			hashMap.Set(195, "foo4")

			timeNow := time.Now()
			removedValue := hashMap.Remove(95)
			fmt.Println(time.Since(timeNow))

			assert.Equal(t, "foo2", *removedValue)

			getValue := hashMap.Get(95)
			assert.Nil(t, getValue)
		})
		t.Run("return nil if key does not exist", func(t *testing.T) {
			hashMap := hashing_solution.NewHashMap[string]()

			hashMap.Set(45, "foo")
			hashMap.Set(10, "foo2")

			removedValue := hashMap.Remove(22)
			assert.Nil(t, removedValue)
		})
	})
}
