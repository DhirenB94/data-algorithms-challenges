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

	if h.Buckets[indexPosition] == nil {
		fmt.Printf("element with key: %v not present\n", key)
		return false
	}
	return true
}
