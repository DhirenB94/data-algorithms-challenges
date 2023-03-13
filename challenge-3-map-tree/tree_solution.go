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
	return searchNode(bst.Root, key)
}

func (bst *BinarySearchTree) Get(key int) (string, bool) {
	return "", false
}

func (bst *BinarySearchTree) Set(key int, value string) string {
	return ""
}

func (bst *BinarySearchTree) Remove(key int) (string, bool) {
	return "", false
}

func searchNode(treeNode *TreeNode, key int) bool {
	//in the case of an empty BST or when you reach a leaf node
	if treeNode == nil {
		return false
	}
	if key == treeNode.Key {
		return true
	}
	//if key is smaller than the key of the current node, keep searching to the left
	if key < treeNode.Key {
		return searchNode(treeNode.LeftNode, key)
	}
	//if key is greater than the key of the current node, keep searching to the Right
	if key > treeNode.Key {
		return searchNode(treeNode.RightNode, key)
	}
	return false
}
