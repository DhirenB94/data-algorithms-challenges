package tree_test

import (
	tree "data/challenge-3-map-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		t.Run("return true when the element is found", func(t *testing.T) {
			rootNode := &tree.TreeNode{
				Key:   50,
				Value: "abc",
			}
			insertTestNode(rootNode, 100, "def")
			insertTestNode(rootNode, 75, "ghi")
			insertTestNode(rootNode, 25, "jkl")
			insertTestNode(rootNode, 1, "mno")

			newBst := tree.NewBinarySearchTree(rootNode)

			shouldContain := newBst.Has(75)
			assert.True(t, shouldContain)
			shouldNotContain := newBst.Has(99)
			assert.False(t, shouldNotContain)
		})
		t.Run("return false when the element is found", func(t *testing.T) {

		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("", func(t *testing.T) {

		})
	})
	t.Run("Set", func(t *testing.T) {
		t.Run("", func(t *testing.T) {

		})
	})
	t.Run("Remove", func(t *testing.T) {
		t.Run("", func(t *testing.T) {

		})
	})
}

func insertTestNode(node *tree.TreeNode, key int, value string) *tree.TreeNode {
	if node == nil {
		newNode := &tree.TreeNode{
			Key:   key,
			Value: value,
		}
		return newNode
	}
	if key < node.Key {
		node.LeftNode = insertTestNode(node.LeftNode, key, value)
	}
	if key > node.Key {
		node.RightNode = insertTestNode(node.RightNode, key, value)
	}
	return node
}
