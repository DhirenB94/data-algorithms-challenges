package main

import (
	"data/challenge-1-map-simple/simple-slice-solution"
	hashing_solution "data/challenge-2-map-hashtable/hashing-solution"
	tree "data/challenge-3-map-tree"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreadSafety(t *testing.T) {
	t.Run("Non thread safe options", func(t *testing.T) {
		t.Run("SimpleKeyValue", func(t *testing.T) {
			//sometimes only is not threadsafe
			simpleKeyValues := simple.NewSliceKeyValues()

			got := threadSafetyChecker(simpleKeyValues)

			for i := 0; i < 1000; i++ {
				//comment print statement out for weird stuff
				fmt.Println("Getting...")
				getValue, _ := got.Get(i)
				assert.Empty(t, getValue)
			}
		})
		t.Run("HashMap", func(t *testing.T) {
			//only works sometimes
			hashMap := hashing_solution.NewHashMap()

			gotHashMap := threadSafetyChecker(hashMap)

			for i := 0; i < 1000; i++ {
				//comment print statement out for weird stuff
				fmt.Println("Getting...")
				getValue, _ := gotHashMap.Get(i)
				assert.Empty(t, getValue)
			}

		})
		t.Run("BinarySearchTree", func(t *testing.T) {
			//sometimes panics, sometimes test passes
			binarySearchTree := tree.NewBinarySearchTree()

			gotTree := threadSafetyChecker(binarySearchTree)

			for i := 0; i < 1000; i++ {
				//comment print statement out for weird stuff
				fmt.Println("Getting...")
				getValue, _ := gotTree.Get(i)
				assert.Empty(t, getValue)
			}
		})
	})
	t.Run("Thread safe options", func(t *testing.T) {
		t.Run("Wrapped SimpleKeyValue", func(t *testing.T) {
			//only works soemetimes
			simpleKeyValues := simple.NewSliceKeyValues()

			threadSafeWrapper := NewThreadSafeMap(simpleKeyValues)
			got := threadSafetyChecker(threadSafeWrapper)

			for i := 0; i < 1000; i++ {
				//comment print statement out for weird stuff
				fmt.Println("Getting...")
				getValue, _ := got.Get(i)
				assert.Empty(t, getValue)
			}

		})
		t.Run("Wrapped HashMap", func(t *testing.T) {
			//only works rarely
			hashMap := hashing_solution.NewHashMap()

			threadSafeWrapper := NewThreadSafeMap(hashMap)
			got := threadSafetyChecker(threadSafeWrapper)

			for i := 0; i < 1000; i++ {
				//comment print statement out for weird stuff
				fmt.Println("Getting...")
				getValue, _ := got.Get(i)
				assert.Empty(t, getValue)
			}
		})
		t.Run("Wrapped BinarySearchTree", func(t *testing.T) {
			//PANICS but works in the main
			bst := tree.NewBinarySearchTree()

			threadSafeWrapper := NewThreadSafeMap(bst)
			got := threadSafetyChecker(threadSafeWrapper)

			for i := 0; i < 1000; i++ {
				//comment print statement out for weird stuff
				fmt.Println("Getting...")
				getValue, _ := got.Get(i)
				assert.Empty(t, getValue)
			}
		})
	})
}
