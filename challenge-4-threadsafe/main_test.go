package main

import (
	"data/challenge-1-map-simple/simple-slice-solution"
	hashing_solution "data/challenge-2-map-hashtable/hashing-solution"
	tree "data/challenge-3-map-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreadSafety(t *testing.T) {
	t.Run("Non thread safe options do not return an empty map", func(t *testing.T) {
		t.Run("SimpleKeyValue", func(t *testing.T) {
			simpleKeyValues := simple.NewSliceKeyValues()
			assert.Empty(t, simpleKeyValues)

			got := ThreadSafetyChecker(simpleKeyValues)

			assert.NotEmpty(t, got)
		})
		t.Run("HashMap", func(t *testing.T) {
			//only works sometimes, dont know how to prove it
			hashMap := hashing_solution.NewHashMap()

			gotHashMap := ThreadSafetyChecker(hashMap)

			for i := 0; i < 1000; i++ {
				getValue, _ := gotHashMap.Get(i)
				assert.Empty(t, getValue)
			}

		})
		t.Run("BinarySearchTree", func(t *testing.T) {
			// Panics
			binarySearchTree := tree.NewBinarySearchTree()

			gotTree := ThreadSafetyChecker(binarySearchTree)

			for i := 0; i < 1000; i++ {
				getValue, _ := gotTree.Get(i)
				assert.Empty(t, getValue)
			}
		})
	})

}
