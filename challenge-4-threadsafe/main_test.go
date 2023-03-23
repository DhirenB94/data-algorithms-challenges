package main

import (
	"data/challenge-1-map-simple/simple-slice-solution"
	hashing_solution "data/challenge-2-map-hashtable/hashing-solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreadSafety(t *testing.T) {
	t.Run("SimpleKeyValue", func(t *testing.T) {
		t.Run("not thread safe", func(t *testing.T) {
			//mostly panics and passes the test this way, but sometimes passes with a non empty map after the checker
			simpleKeyValues := simple.NewSliceKeyValues()

			got := threadSafetyChecker(simpleKeyValues)

			for i := 0; i < 1000; i++ {
				_, isGot := got.Get(i)
				if isGot {
					return
				}
			}
			assert.Fail(t, "should not get here")
		})
		t.Run("is thread safe", func(t *testing.T) {
			simpleKeyValues := simple.NewSliceKeyValues()
			threadSafeWrapper := NewThreadSafeMap(simpleKeyValues)

			got := threadSafetyChecker(threadSafeWrapper)

			for i := 0; i < 1000; i++ {
				getValue, _ := got.Get(i)
				assert.Empty(t, getValue)
			}
		})
	})
	t.Run("Hash", func(t *testing.T) {
		t.Run("not thread safe", func(t *testing.T) {
			//only sometimes is it not threadsafe - ie the test passes
			hashmap := hashing_solution.NewHashMap()

			got := threadSafetyChecker(hashmap)

			for i := 0; i < 1000; i++ {
				_, isGot := got.Get(i)
				if isGot {
					return
				}
			}
			assert.Fail(t, "should not get here")
		})
		t.Run("is thread safe", func(t *testing.T) {
			hashmap := hashing_solution.NewHashMap()
			wrapper := NewThreadSafeMap(hashmap)

			got := threadSafetyChecker(wrapper)

			for i := 0; i < 1000; i++ {
				getValue, _ := got.Get(i)
				assert.Empty(t, getValue)
			}
		})
	})
}
