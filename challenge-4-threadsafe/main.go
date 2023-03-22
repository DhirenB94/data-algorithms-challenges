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

type threadSafeMap struct {
	mutex     sync.Mutex
	customMap interfaces.Operations
}

func NewThreadSafeMap(customMap interfaces.Operations) interfaces.Operations {
	return &threadSafeMap{customMap: customMap}
}

func (tsm *threadSafeMap) Has(key int) bool {
	tsm.mutex.Lock()
	defer tsm.mutex.Unlock()
	return tsm.customMap.Has(key)
}

func (tsm *threadSafeMap) Get(key int) (string, bool) {
	tsm.mutex.Lock()
	defer tsm.mutex.Unlock()
	return tsm.customMap.Get(key)
}
func (tsm *threadSafeMap) Remove(key int) (string, bool) {
	tsm.mutex.Lock()
	defer tsm.mutex.Unlock()
	return tsm.customMap.Remove(key)
}
func (tsm *threadSafeMap) Set(key int, value string) (string, bool) {
	tsm.mutex.Lock()
	defer tsm.mutex.Unlock()
	return tsm.customMap.Set(key, value)
}

func main() {
	simpleKeyValues := simple.NewSliceKeyValues()
	//threadSafetyChecker(simpleKeyValues)
	simpleKeyValuesWrapper := NewThreadSafeMap(simpleKeyValues)
	threadSafetyChecker(simpleKeyValuesWrapper)

	hashMap := hashing_solution.NewHashMap()
	//threadSafetyChecker(hashMap)
	hashMapWrapper := NewThreadSafeMap(hashMap)
	threadSafetyChecker(hashMapWrapper)

	binarySearchTree := tree.NewBinarySearchTree()
	//threadSafetyChecker(binarySearchTree)
	binarySearchTreeWrapper := NewThreadSafeMap(binarySearchTree)
	threadSafetyChecker(binarySearchTreeWrapper)

}

func threadSafetyChecker(anyMap interfaces.Operations) interfaces.Operations {
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
			fmt.Printf("Goroutine 1 SET: setting key = %v\n", i)
			wg.Done()
		}
	}()

	go func() {
		//defer wg.Done()
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			removedValues, _ := anyMap.Remove(i)
			fmt.Printf("Goroutine 2 REMOVE: removedValue = %v, key = %v\n", removedValues, i)
			wg.Done()
		}
	}()

	go func() {
		//defer wg.Done()
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			value, _ := anyMap.Get(i)
			fmt.Printf("Goroutine 3 GET: gotValue = %v, key = %v\n", value, i)
			wg.Done()
		}
	}()

	// Wait for the 3 goroutines to complete
	wg.Wait()

	fmt.Println(anyMap)
	return anyMap
}
