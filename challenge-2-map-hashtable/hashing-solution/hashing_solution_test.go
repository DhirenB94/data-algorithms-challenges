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
			hashTable := newHashMap(50)
			newNode := &hashing_solution.Node{
				Key:   45,
				Value: "foo",
			}
			nextNode := &hashing_solution.Node{
				Key:   95,
				Value: "bar",
			}

			hashTable.Buckets[45] = newNode
			newNode.Next = nextNode

			timeNow := time.Now()
			answer := hashTable.Has(45)
			fmt.Println(time.Since(timeNow))

			assert.Equal(t, true, answer)

			answer2 := hashTable.Has(95)
			assert.Equal(t, true, answer2)
		})
		t.Run("return false if key is not present", func(t *testing.T) {
			hashTable := newHashMap(50)

			answer := hashTable.Has(10)
			assert.Equal(t, false, answer)
		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("gets the value if key present", func(t *testing.T) {
			hashTable := newHashMap(50)
			newNode := &hashing_solution.Node{
				Key:   45,
				Value: "foo",
			}
			nextNode := &hashing_solution.Node{
				Key:   95,
				Value: "bar",
			}

			hashTable.Buckets[45] = newNode
			newNode.Next = nextNode

			timeNow := time.Now()
			answer, _ := hashTable.Get(45)
			fmt.Println(time.Since(timeNow))

			assert.Equal(t, "foo", answer)

			answer2, _ := hashTable.Get(95)
			assert.Equal(t, "bar", answer2)
		})
		t.Run("returns empty string and false if key not present", func(t *testing.T) {
			hashTable := newHashMap(50)

			answer, isPresent := hashTable.Get(45)
			assert.Equal(t, "", answer)
			assert.False(t, isPresent)
		})
	})
	t.Run("Set", func(t *testing.T) {
		t.Run("adds the element if doesnt exist already", func(t *testing.T) {
			hashMap := newHashMap(50)

			timeNow := time.Now()
			setValue, isSet := hashMap.Set(45, "foo")
			fmt.Println(time.Since(timeNow))

			assert.Empty(t, setValue)
			assert.True(t, isSet)

			getValue, _ := hashMap.Get(45)
			assert.Equal(t, "foo", getValue)
			assert.Equal(t, getValue, hashMap.Buckets[45].Value)

			setValue2, _ := hashMap.Set(95, "bar")
			assert.Empty(t, setValue2)

			getValue2, _ := hashMap.Get(95)
			assert.Equal(t, "bar", getValue2)
			assert.Equal(t, getValue2, hashMap.Buckets[45].Next.Value)

		})
		t.Run("if element with key exists already, update its value", func(t *testing.T) {
			hashMap := newHashMap(50)

			setValue, _ := hashMap.Set(45, "foo")
			assert.Empty(t, setValue)

			timeNow := time.Now()
			oldValue, isSet := hashMap.Set(45, "bar")
			fmt.Println(time.Since(timeNow))

			assert.Equal(t, "foo", oldValue)
			assert.True(t, isSet)

			getValue, _ := hashMap.Get(45)
			assert.Equal(t, "bar", getValue)
		})
	})
	t.Run("Remove", func(t *testing.T) {
		t.Run("remove key if exists at the head of linked list", func(t *testing.T) {
			hashMap := newHashMap(50)

			hashMap.Set(45, "foo")
			hashMap.Set(95, "foo2")
			hashMap.Set(145, "foo3")

			answer, isPresent := hashMap.Remove(45)
			assert.Equal(t, "foo", answer)
			assert.True(t, isPresent)

			getValue, isPresent := hashMap.Get(45)
			assert.Empty(t, getValue)
			assert.False(t, isPresent)
		})
		t.Run("remove key if exists somewhere in the linked list", func(t *testing.T) {
			hashMap := newHashMap(50)
			hashMap.Set(5, "foo")
			hashMap.Set(55, "bar")
			hashMap.Set(105, "cho")
			hashMap.Set(155, "mod")

			timeNow := time.Now()
			oldValue, isPresent := hashMap.Remove(55)
			fmt.Println(time.Since(timeNow))

			assert.Equal(t, "bar", oldValue)
			assert.True(t, isPresent)

			getValue, isPresent2 := hashMap.Get(55)
			assert.Empty(t, getValue)
			assert.False(t, isPresent2)
		})
		t.Run("return empty string and false if key does not exist", func(t *testing.T) {
			hashMap := newHashMap(50)

			answer, isPresent := hashMap.Remove(45)
			assert.Empty(t, answer)
			assert.False(t, isPresent)

		})
	})
}

func newHashMap(size int) *hashing_solution.HashMap {
	return &hashing_solution.HashMap{
		Buckets: make([]*hashing_solution.Node, size),
	}
}
