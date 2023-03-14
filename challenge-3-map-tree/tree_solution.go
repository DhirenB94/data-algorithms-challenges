package tree

type TreeNode struct {
	Key       int
	Value     string
	LeftNode  *TreeNode
	RightNode *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

func NewBinarySearchTree(treeNode *TreeNode) *BinarySearchTree {
	return &BinarySearchTree{Root: treeNode}
}

func (bst *BinarySearchTree) Has(key int) bool {
	_, isPresent := searchNode(bst.Root, key)
	return isPresent
}

func (bst *BinarySearchTree) Get(key int) (string, bool) {
	return searchNode(bst.Root, key)
}

func (bst *BinarySearchTree) Set(key int, value string) (string, bool) {
	return insertNode(bst.Root, key, value)
}

func (bst *BinarySearchTree) Remove(key int) (string, bool) {
	return "", false
}

func searchNode(treeNode *TreeNode, key int) (string, bool) {
	//in the case of an empty BST or when you reach a leaf node
	if treeNode == nil {
		return "", false
	}
	if key == treeNode.Key {
		return treeNode.Value, true
	}
	//if key is smaller than the key of the current node, keep searching to the left
	if key < treeNode.Key {
		return searchNode(treeNode.LeftNode, key)
	}
	//if key is greater than the key of the current node, keep searching to the Right
	if key > treeNode.Key {
		return searchNode(treeNode.RightNode, key)
	}
	return "", false
}

func insertNode(node *TreeNode, key int, value string) (string, bool) {
	newNode := &TreeNode{
		Key:   key,
		Value: value,
	}
	//if there's no root node or reached a leaf node, insert the node
	if node == nil {
		node = newNode
		return "", true
	}
	//If the key is present update the value and return the old value
	if key == node.Key {
		oldValue := node.Value
		node.Value = value
		return oldValue, true
	}
	if key < node.Key {
		if node.LeftNode == nil {
			node.LeftNode = newNode
			return "", true
		}
		return insertNode(node.LeftNode, key, value)
	}
	if key > node.Key {
		if node.RightNode == nil {
			node.RightNode = newNode
			return "", true
		}
		return insertNode(node.RightNode, key, value)
	}
	return "", false
}
