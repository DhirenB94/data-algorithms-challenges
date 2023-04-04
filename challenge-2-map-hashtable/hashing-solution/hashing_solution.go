package hashing_solution

import (
	interfaces "data"
	"fmt"
)

type node[T any] struct {
	key   int
	value T
	next  *node[T]
}

type hashMap[T any] struct {
	buckets []*node[T]
}

func NewHashMap[T any]() interfaces.Operations[T] {
	return &hashMap[T]{buckets: make([]*node[T], 50)}
}

func hash(key int) int {
	return key % 50
}

func (h *hashMap[T]) Has(key int) bool {
	indexPosition := hash(key)
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

func (h *hashMap[T]) Get(key int) *T {
	indexPosition := hash(key)
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

func (h *hashMap[T]) Set(key int, value T) *T {
	index := hash(key)
	newNode := &node[T]{
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

func (h *hashMap[T]) Remove(key int) *T {
	indexPosition := hash(key)
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

func (h *hashMap[T]) String() string {
	return fmt.Sprint("Bucket", h.buckets)
}
