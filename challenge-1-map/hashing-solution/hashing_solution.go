package hashing_solution

import (
	"fmt"
)

type Node struct {
	Key   int
	Value string
	Next  *Node
}

type HashMap struct {
	Buckets []*Node
}

func hash(key int) int {
	return key % 50
}

func (h *HashMap) Has(key int) bool {
	indexPosition := hash(key)
	if h.Buckets[indexPosition] != nil {
		//iterate through all elements in the list at that index position
		firstNodeInList := h.Buckets[indexPosition]
		for n := firstNodeInList; n != nil; n = n.Next {
			if n.Key == key {
				return true
			}
		}
	}
	fmt.Printf("element with key: %v not present\n", key)
	return false
}

func (h *HashMap) Get(key int) (string, bool) {
	indexPosition := hash(key)
	if h.Buckets[indexPosition] != nil {
		//check if it is in the linked list of the index position
		firstNodeInList := h.Buckets[indexPosition]
		for n := firstNodeInList; n != nil; n = n.Next {
			if n.Key == key {
				return n.Value, true
			}
		}
	}
	fmt.Printf("element with key: %v not present\n", key)
	return "", false
}

func (h *HashMap) Set(key int, value string) string {
	index := hash(key)
	newNode := &Node{
		Key:   key,
		Value: value,
	}
	//if index is empty, can insert the node
	if h.Buckets[index] == nil {
		h.Buckets[index] = newNode
	}
	//if some nodes already exists at that index position, loop through the linked-list
	firstNodeInList := h.Buckets[index]
	for n := firstNodeInList; n != nil; n = n.Next {
		//if the key exists within the list, update its value
		if n.Key == key {
			n.Value = value
			return value
		}
		//if the key does not exist, add to the end of the linked list
		firstNodeInList.Next = newNode
	}
	return value
}
