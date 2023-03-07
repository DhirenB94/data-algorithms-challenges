package hashing_solution_test

import (
	hashing_solution "data/challenge-1-map/hashing-solution"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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

			answer := hashTable.Has(45)
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

			answer, _ := hashTable.Get(45)
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

			newValue := hashMap.Set(45, "foo")
			getValue, _ := hashMap.Get(45)
			assert.Equal(t, newValue, hashMap.Buckets[45].Value)
			assert.Equal(t, newValue, getValue)

			newValue2 := hashMap.Set(95, "bar")
			getValue2, _ := hashMap.Get(95)
			assert.Equal(t, newValue2, hashMap.Buckets[45].Next.Value)
			assert.Equal(t, newValue2, getValue2)

		})
		t.Run("if element with key exists already, update its value", func(t *testing.T) {
			hashMap := newHashMap(50)

			oldValue := hashMap.Set(45, "foo")
			newValue := hashMap.Set(45, "bar")

			getValue, _ := hashMap.Get(45)
			fmt.Println(getValue)

			assert.Equal(t, newValue, getValue)
			assert.NotEqual(t, oldValue, getValue)
		})
	})
}

func newHashMap(size int) *hashing_solution.HashMap {
	return &hashing_solution.HashMap{
		Buckets: make([]*hashing_solution.Node, size),
	}
}
