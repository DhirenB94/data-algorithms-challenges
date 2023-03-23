package tree

import (
	interfaces "data"
	"fmt"
)

type treeNode struct {
	key       int
	value     string
	leftNode  *treeNode
	rightNode *treeNode
}

type binarySearchTree struct {
	root *treeNode
}

func NewBinarySearchTree() interfaces.Operations {
	return &binarySearchTree{root: &treeNode{}}
}

func (bst *binarySearchTree) Has(key int) bool {
	_, isPresent := searchNode(bst.root, key)
	return isPresent
}

func (bst *binarySearchTree) Get(key int) (string, bool) {
	return searchNode(bst.root, key)
}

func (bst *binarySearchTree) Set(key int, value string) (string, bool) {
	newNode := &treeNode{
		key:   key,
		value: value,
	}
	if bst.root == nil {
		bst.root = newNode
		return "", true
	}
	return insertNode(bst.root, key, value)
}

func (bst *binarySearchTree) Remove(key int) (string, bool) {
	getValue, isGot := bst.Get(key)
	if !isGot {
		return "", false
	}
	removeNode(bst.root, key)
	return getValue, true
}

func (bst *binarySearchTree) String() string {
	return fmt.Sprint("Root", bst.root)
}

func searchNode(treeNode *treeNode, key int) (string, bool) {
	//in the case of an empty BST or when you reach a leaf node
	if treeNode == nil {
		return "", false
	}
	if key == treeNode.key {
		return treeNode.value, true
	}
	//if key is smaller than the key of the current node, keep searching to the left
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}
	//if key is greater than the key of the current node, keep searching to the Right
	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return "", false
}

func insertNode(node *treeNode, key int, value string) (string, bool) {
	newNode := &treeNode{
		key:   key,
		value: value,
	}

	//If the key is present update the value and return the old value
	if key == node.key {
		oldValue := node.value
		node.value = value
		return oldValue, true
	}
	if key < node.key {
		if node.leftNode == nil {
			node.leftNode = newNode
			return "", true
		}
		return insertNode(node.leftNode, key, value)
	}
	if key > node.key {
		if node.rightNode == nil {
			node.rightNode = newNode
			return "", true
		}
		return insertNode(node.rightNode, key, value)
	}
	return "", false
}

func removeNode(node *treeNode, key int) *treeNode {
	if node == nil {
		return nil
	}
	//Finding the key
	if key < node.key {
		node.leftNode = removeNode(node.leftNode, key)
	}
	if key > node.key {
		node.rightNode = removeNode(node.rightNode, key)
	}
	if key == node.key {
		//if the node to be removed is a leaf node
		if node.leftNode == nil && node.rightNode == nil {
			node = nil
			return nil
		}
		//if the node to be removed has a L/R subtree only
		if node.leftNode == nil {
			node = node.rightNode
		}
		if node.rightNode == nil {
			node = node.leftNode
		}
		//if the node to be removed has BOTH subtrees
		if node.leftNode != nil && node.rightNode != nil {
			//chose the R subtree and go as far left as possible to find the smallest node
			smallestNode := smallestNode(node.rightNode)
			//copy the smallest node found to the node to be removed
			node.key = smallestNode.key
			node.value = smallestNode.value
			//the smallest node will always either be a leaf node or have a max of 1 child (in this case 1 right child only)
			//call the remove method on it, but start from the right node of the original (as youve copied smallest node to the original, it would re-enter the loop)
			node.rightNode = removeNode(node.rightNode, smallestNode.key)
		}
	}
	return node
}

func smallestNode(node *treeNode) *treeNode {
	for node.leftNode != nil {
		node = node.leftNode
	}
	return node
}
