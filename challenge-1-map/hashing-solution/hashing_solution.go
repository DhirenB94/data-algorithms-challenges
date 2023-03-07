package hashing_solution

import "fmt"

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

func (h HashMap) Get(key int) (string, bool) {
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
