package threadsafe

import (
	interfaces "data/common-interfaces"
	"fmt"
	"sync"
)

type threadSafeMap[K, V interfaces.CustomType] struct {
	mutex     sync.Mutex
	customMap interfaces.Operations[K, V]
}

func NewThreadSafeMap[K, V interfaces.CustomType](customMap interfaces.Operations[K, V]) interfaces.Operations[K, V] {
	return &threadSafeMap[K, V]{customMap: customMap[K, V]}
}

func (tsm *threadSafeMap[K, V]) Has(key K) bool {
	tsm.mutex.Lock()
	defer tsm.mutex.Unlock()
	return tsm.customMap.Has(key)
}

func (tsm *threadSafeMap[K, V]) Get(key K) *V {
	tsm.mutex.Lock()
	defer tsm.mutex.Unlock()
	return tsm.customMap.Get(key)
}
func (tsm *threadSafeMap[K, V]) Remove(key K) *V {
	tsm.mutex.Lock()
	defer tsm.mutex.Unlock()
	return tsm.customMap.Remove(key)
}
func (tsm *threadSafeMap[K, V]) Set(key K, value V) *V {
	tsm.mutex.Lock()
	defer tsm.mutex.Unlock()
	return tsm.customMap.Set(key, value)
}
func (tsm *threadSafeMap[K, V]) String() string {
	return fmt.Sprint("custom map", tsm.customMap)
}

//func ThreadSafetyChecker[K, V interfaces.CustomType](anyMap interfaces.Operations[K, V]) interfaces.Operations[K, V] {
//	var wg sync.WaitGroup
//
//	for worker := 0; worker < 100; worker++ {
//		wg.Add(1)
//		go func(worker int) {
//			for i := 0; i < 100; i++ {
//				key := worker*100 + i
//				value := strconv.Itoa(key)
//				anyMap.Set(K(key), value)
//			}
//			for i := 0; i < 100; i++ {
//				key := worker*100 + i
//				anyMap.Get(K(key))
//			}
//			for i := 0; i < 100; i++ {
//				key := worker*100 + i
//				anyMap.Remove(K(key))
//			}
//			wg.Done()
//		}(worker)
//	}
//	wg.Wait()
//
//	fmt.Println(anyMap)
//	return anyMap
//}
