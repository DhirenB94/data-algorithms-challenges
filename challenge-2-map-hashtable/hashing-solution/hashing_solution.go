package hashing_solution

import (
	interfaces "data"
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

func NewHashMap() interfaces.Operations {
	return &HashMap{Buckets: make([]*Node, 50)}
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

func (h *HashMap) Set(key int, value string) (string, bool) {
	index := hash(key)
	newNode := &Node{
		Key:   key,
		Value: value,
	}
	//if index is empty, can insert the node
	if h.Buckets[index] == nil {
		h.Buckets[index] = newNode
		return "", true
	}
	//if some nodes already exists at that index position, loop through the linked-list
	firstNodeInList := h.Buckets[index]
	for n := firstNodeInList; n != nil; n = n.Next {
		//if the key exists within the list, update its value
		if n.Key == key {
			oldValue := n.Value
			n.Value = value
			return oldValue, true
		}
		// If it doesnt exist, add the new node to the end of the list
		if n.Next == nil {
			n.Next = newNode
			return "", true
		}
	}
	return "", false
}

func (h *HashMap) Remove(key int) (string, bool) {
	indexPosition := hash(key)
	oldValue, _ := h.Get(key)
	//handle when you need to remove the 1st node in the list
	if h.Buckets[indexPosition] != nil && h.Buckets[indexPosition].Key == key {
		h.Buckets[indexPosition] = h.Buckets[indexPosition].Next
		return oldValue, true
	}
	//handle when you need to remove nth node in a list
	for n := h.Buckets[indexPosition]; n != nil; n = n.Next {
		if n.Next != nil && n.Next.Key == key {
			n.Next = n.Next.Next
			return oldValue, true
		}
	}
	return "", false
}
