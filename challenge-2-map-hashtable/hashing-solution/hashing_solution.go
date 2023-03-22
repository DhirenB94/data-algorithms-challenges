package hashing_solution

import (
	interfaces "data"
)

type node struct {
	key   int
	value string
	next  *node
}

type hashMap struct {
	buckets []*node
}

func NewHashMap() interfaces.Operations {
	return &hashMap{buckets: make([]*node, 50)}
}

func hash(key int) int {
	return key % 50
}

func (h *hashMap) Has(key int) bool {
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

func (h *hashMap) Get(key int) (string, bool) {
	indexPosition := hash(key)
	if h.buckets[indexPosition] != nil {
		//check if it is in the linked list of the index position
		firstNodeInList := h.buckets[indexPosition]
		for n := firstNodeInList; n != nil; n = n.next {
			if n.key == key {
				return n.value, true
			}
		}
	}
	return "", false
}

func (h *hashMap) Set(key int, value string) (string, bool) {
	index := hash(key)
	newNode := &node{
		key:   key,
		value: value,
	}
	//if index is empty, can insert the node
	if h.buckets[index] == nil {
		h.buckets[index] = newNode
		return "", true
	}
	//if some nodes already exists at that index position, loop through the linked-list
	firstNodeInList := h.buckets[index]
	for n := firstNodeInList; n != nil; n = n.next {
		//if the key exists within the list, update its value
		if n.key == key {
			oldValue := n.value
			n.value = value
			return oldValue, true
		}
		// If it doesnt exist, add the new node to the end of the list
		if n.next == nil {
			n.next = newNode
			return "", true
		}
	}
	return "", false
}

func (h *hashMap) Remove(key int) (string, bool) {
	indexPosition := hash(key)
	oldValue, _ := h.Get(key)
	//handle when you need to remove the 1st node in the list
	if h.buckets[indexPosition] != nil && h.buckets[indexPosition].key == key {
		h.buckets[indexPosition] = h.buckets[indexPosition].next
		return oldValue, true
	}
	//handle when you need to remove nth node in a list
	for n := h.buckets[indexPosition]; n != nil; n = n.next {
		if n.next != nil && n.next.key == key {
			n.next = n.next.next
			return oldValue, true
		}
	}
	return "", false
}
