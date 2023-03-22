package main

import (
	interfaces "data"
	"data/challenge-1-map-simple/simple-slice-solution"
	hashing_solution "data/challenge-2-map-hashtable/hashing-solution"
	tree "data/challenge-3-map-tree"
	"fmt"
	"strconv"
	"sync"
)

func main() {
	simpleKeyValues := simple.NewSliceKeyValues()
	hashMap := hashing_solution.NewHashMap()
	binarySearchTree := tree.NewBinarySearchTree()

	ThreadSafetyChecker(simpleKeyValues)

	ThreadSafetyChecker(hashMap)

	ThreadSafetyChecker(binarySearchTree)

}

func ThreadSafetyChecker(customMap interfaces.Operations) interfaces.Operations {
	fmt.Println(customMap)

	// Create three goroutines that concurrently access the customMap instance
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			value := strconv.Itoa(i)
			customMap.Set(i, value)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			removedValues, _ := customMap.Remove(i)
			fmt.Printf("Goroutine 2 REMOVE: removedValue = %v, key = %v\n", removedValues, i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			value, _ := customMap.Get(i)
			fmt.Printf("Goroutine 2 GET: gotValue = %v, key = %v\n", value, i)
		}
	}()

	// Wait for the 3 goroutines to complete
	wg.Wait()

	fmt.Println(customMap)
	return customMap
}
