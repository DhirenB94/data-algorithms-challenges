package hashing_solution

import (
	interfaces "data"
	"fmt"
)

type node[K, V interfaces.CustomType] struct {
	key   K
	value V
	next  *node[K, V]
}

type hashMap[K, V interfaces.CustomType] struct {
	buckets []*node[K, V]
}

func NewHashMap[K, V interfaces.CustomType]() interfaces.Operations[K, V] {
	return &hashMap[K, V]{buckets: make([]*node[K, V], 50)}
}

func (h *hashMap[K, V]) Has(key K) bool {
	indexPosition := key.HashMe()
	if h.buckets[indexPosition] != nil {
		//iterate through all elements in the list at that index position
		firstNodeInList := h.buckets[indexPosition]
		for n := firstNodeInList; n != nil; n = n.next {
			if n.key == key {
				return true
			}
		}
	}
	return false
}

func (h *hashMap[K, V]) Get(key K) *V {
	indexPosition := key.HashMe()
	if h.buckets[indexPosition] != nil {
		//check if it is in the linked list of the index position
		firstNodeInList := h.buckets[indexPosition]
		for n := firstNodeInList; n != nil; n = n.next {
			if n.key == key {
				return &n.value
			}
		}
	}
	return nil
}

func (h *hashMap[K, V]) Set(key K, value V) *V {
	index := key.HashMe()
	newNode := &node[K, V]{
		key:   key,
		value: value,
	}
	//if index is empty, can insert the node
	if h.buckets[index] == nil {
		h.buckets[index] = newNode
		return nil
	}
	//if some nodes already exists at that index position, loop through the linked-list
	firstNodeInList := h.buckets[index]
	for n := firstNodeInList; n != nil; n = n.next {
		//if the key exists within the list, update its value
		if n.key == key {
			oldValue := n.value
			n.value = value
			return &oldValue
		}
		// If it doesnt exist, add the new node to the end of the list
		if n.next == nil {
			n.next = newNode
			return nil
		}
	}
	return nil
}

func (h *hashMap[K, V]) Remove(key K) *V {
	indexPosition := key.HashMe()
	oldValue := h.Get(key)
	//handle when you need to remove the 1st node in the list
	if h.buckets[indexPosition] != nil && h.buckets[indexPosition].key == key {
		h.buckets[indexPosition] = h.buckets[indexPosition].next
		return oldValue
	}
	//handle when you need to remove nth node in a list
	for n := h.buckets[indexPosition]; n != nil; n = n.next {
		if n.next != nil && n.next.key == key {
			n.next = n.next.next
			return oldValue
		}
	}
	return nil
}

func (h *hashMap[K, V]) String() string {
	return fmt.Sprint("Bucket", h.buckets)
}
