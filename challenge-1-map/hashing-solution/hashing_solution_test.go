package hashing_solution_test

import (
	hashing_solution "data/challenge-1-map/hashing-solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashKeys(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		t.Run("return true if key is present", func(t *testing.T) {
			hashTable := newHashMap(50)
			newNode := &hashing_solution.Node{
				Key:   5,
				Value: "five",
			}
			hashTable.Buckets[5] = newNode

			answer := hashTable.Has(5)
			assert.Equal(t, true, answer)
		})
		t.Run("return false if key is not present", func(t *testing.T) {
			hashTable := newHashMap(50)

			answer := hashTable.Has(10)
			assert.Equal(t, false, answer)
		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("gets the value if key present", func(t *testing.T) {

		})
		t.Run("returns empty string if key not present", func(t *testing.T) {

		})
	})
}

func newHashMap(size int) *hashing_solution.HashMap {
	return &hashing_solution.HashMap{
		Buckets: make([]*hashing_solution.Node, size),
	}
}
