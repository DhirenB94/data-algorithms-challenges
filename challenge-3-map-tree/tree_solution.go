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
	newNode := &TreeNode{
		Key:   key,
		Value: value,
	}
	if bst.Root == nil {
		bst.Root = newNode
		return "", true
	}
	return insertNode(bst.Root, key, value)
}

func (bst *BinarySearchTree) Remove(key int) (string, bool) {
	getValue, isGot := bst.Get(key)
	if !isGot {
		return "", false
	}
	removeNode(bst.Root, key)
	return getValue, true
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

func removeNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}
	//Finding the key
	if key < node.Key {
		node.LeftNode = removeNode(node.LeftNode, key)
	}
	if key > node.Key {
		node.RightNode = removeNode(node.RightNode, key)
	}
	if key == node.Key {
		//if the node to be removed is a leaf node
		if node.LeftNode == nil && node.RightNode == nil {
			node = nil
			return nil
		}
		//if the node to be removed has a L/R subtree only
		if node.LeftNode == nil {
			node = node.RightNode
		}
		if node.RightNode == nil {
			node = node.LeftNode
		}
		//if the node to be removed has BOTH subtrees
		if node.LeftNode != nil && node.RightNode != nil {
			//chose the R subtree and go as far left as possible to find the smallest node
			smallestNode := smallestNode(node.RightNode)
			//copy the smallest node found to the node to be removed
			node.Key = smallestNode.Key
			node.Value = smallestNode.Value
			//the smallest node will always either be a leaf node or have a max of 1 child (in this case 1 right child only)
			//call the remove method on it, but start from the right node of the original (as youve copied smallest node to the original, it would re-enter the loop)
			removeNode(node.RightNode, smallestNode.Key)
		}
	}
	return node
}

func smallestNode(node *TreeNode) *TreeNode {
	for node.LeftNode != nil {
		node = node.LeftNode
	}
	return node
}
