package tree

import (
	interfaces "data/common-interfaces"
	"fmt"
)

type treeNode[K, V interfaces.CustomType] struct {
	key       K
	value     V
	leftNode  *treeNode[K, V]
	rightNode *treeNode[K, V]
}

type binarySearchTree[K, V interfaces.CustomType] struct {
	root *treeNode[K, V]
}

func NewBinarySearchTree[K, V interfaces.CustomType]() interfaces.Operations[K, V] {
	return &binarySearchTree[K, V]{root: &treeNode[K, V]{}}
}

func (bst *binarySearchTree[K, V]) Has(key K) bool {
	_, isPresent := searchNode(bst.root, key)
	return isPresent
}

func (bst *binarySearchTree[K, V]) Get(key K) *V {
	gotNode, _ := searchNode(bst.root, key)
	return gotNode
}

func (bst *binarySearchTree[K, V]) Set(key K, value V) *V {
	newNode := &treeNode[K, V]{
		key:   key,
		value: value,
	}
	if bst.root == nil {
		bst.root = newNode
		return nil
	}
	return insertNode(bst.root, key, value)
}

func (bst *binarySearchTree[K, V]) Remove(key K) *V {
	return removeNode(bst.root, key)
}

func (bst *binarySearchTree[K, V]) String() string {
	return fmt.Sprint("Root", bst.root)
}

func searchNode[K, V interfaces.CustomType](treeNode *treeNode[K, V], key K) (*V, bool) {
	//in the case of an empty BST or when you reach a leaf node
	if treeNode == nil {
		return nil, false
	}
	if key == treeNode.key {
		return &treeNode.value, true
	}
	//if key is smaller than the key of the current node, keep searching to the left
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}
	//if key is greater than the key of the current node, keep searching to the Right
	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return nil, false
}

func insertNode[K, V interfaces.CustomType](node *treeNode[K, V], key K, value V) *V {
	newNode := &treeNode[K, V]{
		key:   key,
		value: value,
	}

	//If the key is present update the value and return the old value
	if key == node.key {
		oldValue := node.value
		node.value = value
		return &oldValue
	}
	if key < node.key {
		if node.leftNode == nil {
			node.leftNode = newNode
			return nil
		}
		return insertNode(node.leftNode, key, value)
	}
	if key > node.key {
		if node.rightNode == nil {
			node.rightNode = newNode
			return nil
		}
		return insertNode(node.rightNode, key, value)
	}
	return nil
}

func removeNode[K, V interfaces.CustomType](node *treeNode[K, V], key K) *V {
	if node == nil {
		return nil
	}
	//Finding the key
	if key < node.key {
		return removeNode(node.leftNode, key)
	}
	if key > node.key {
		return removeNode(node.rightNode, key)
	}
	if key == node.key {
		//if the node to be removed is a leaf node
		if node.leftNode == nil && node.rightNode == nil {
			removedValue := node.value
			node = nil
			return &removedValue
		}
		//if the node to be removed has a L/R subtree only
		if node.leftNode == nil {
			removedValue := node.value
			node = node.rightNode
			return &removedValue
		}
		if node.rightNode == nil {
			removedValue := node.value
			node = node.leftNode
			return &removedValue
		}
		//if the node to be removed has BOTH subtrees
		if node.leftNode != nil && node.rightNode != nil {
			//chose the R subtree and go as far left as possible to find the smallest node
			smallestNode := smallestNode(node.rightNode)
			//copy the smallest node found to the node to be removed
			node.key = smallestNode.key
			node.value = smallestNode.value
			//remove the smallest node from the right subtree
			removedValue := removeNode(node.rightNode, smallestNode.key)
			return removedValue
		}
	}
	return nil
}

func smallestNode[K, V interfaces.CustomType](node *treeNode[K, V]) *treeNode[K, V] {
	for node.leftNode != nil {
		node = node.leftNode
	}
	return node
}
