package main

import (
	interfaces "data"
	"data/challenge-1-map-simple/simple-slice-solution"
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
func (tsm *threadSafeMap) String() string {
	return fmt.Sprint("custom map", tsm.customMap)
}

func main() {
	simpleKeyValues := simple.NewSliceKeyValues()
	threadSafetyChecker(simpleKeyValues)
	simpleKeyValuesWrapper := NewThreadSafeMap(simpleKeyValues)
	threadSafetyChecker(simpleKeyValuesWrapper)

	//hashMap := hashing_solution.NewHashMap()
	//threadSafetyChecker(hashMap)
	//hashMapWrapper := NewThreadSafeMap(hashMap)
	//threadSafetyChecker(hashMapWrapper)

}

func threadSafetyChecker(anyMap interfaces.Operations) interfaces.Operations {
	var wg sync.WaitGroup

	for worker := 0; worker < 100; worker++ {
		wg.Add(1)
		go func(worker int) {
			for i := 0; i < 100; i++ {
				key := worker*100 + i
				value := strconv.Itoa(key)
				anyMap.Set(key, value)
			}
			for i := 0; i < 100; i++ {
				key := worker*100 + i
				anyMap.Get(key)
			}
			for i := 0; i < 100; i++ {
				key := worker*100 + i
				anyMap.Remove(key)
			}
			wg.Done()
		}(worker)
	}
	wg.Wait()

	fmt.Println(anyMap)
	return anyMap
}
