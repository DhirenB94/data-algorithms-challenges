package main

import (
	interfaces "data"
	tree "data/challenge-3-map-tree"
	"fmt"
	"strconv"
	"sync"
)

func main() {
	//simpleKeyValues := simple.NewSliceKeyValues()
	//ThreadSafetyChecker(simpleKeyValues)

	//hashMap := hashing_solution.NewHashMap()
	//ThreadSafetyChecker(hashMap)
	//
	binarySearchTree := tree.NewBinarySearchTree()
	ThreadSafetyChecker(binarySearchTree)

}

func ThreadSafetyChecker(anyMap interfaces.Operations) interfaces.Operations {
	fmt.Println(anyMap)

	var wg sync.WaitGroup
	//wg.Add(3)

	// Create three goroutines that concurrently access the anyMap instance
	go func() {
		//defer wg.Done()
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			value := strconv.Itoa(i)
			anyMap.Set(i, value)
		}
	}()

	go func() {
		//defer wg.Done()
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			removedValues, _ := anyMap.Remove(i)
			fmt.Printf("Goroutine 2 REMOVE: removedValue = %v, key = %v\n", removedValues, i)
		}
	}()

	go func() {
		//defer wg.Done()
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			value, _ := anyMap.Get(i)
			fmt.Printf("Goroutine 2 GET: gotValue = %v, key = %v\n", value, i)
		}
	}()

	// Wait for the 3 goroutines to complete
	wg.Wait()

	fmt.Println(anyMap)
	return anyMap
}
