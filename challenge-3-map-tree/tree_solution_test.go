package tree_test

import (
	tree "data/challenge-3-map-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		t.Run("return true when the element is found and false when it is not", func(t *testing.T) {
			rootNode := &tree.TreeNode{
				Key:   50,
				Value: "abc",
			}

			newBst := tree.NewBinarySearchTree(rootNode)

			newBst.Set(100, "def")
			newBst.Set(75, "ghi")
			newBst.Set(25, "jkl")
			newBst.Set(1, "mno")

			shouldContain := newBst.Has(75)
			assert.True(t, shouldContain)
			shouldNotContain := newBst.Has(99)
			assert.False(t, shouldNotContain)
		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("return the value and true when the element is found", func(t *testing.T) {
			rootNode := &tree.TreeNode{
				Key:   50,
				Value: "abc",
			}

			newBst := tree.NewBinarySearchTree(rootNode)

			newBst.Set(100, "def")
			newBst.Set(25, "jkl")
			newBst.Set(10, "")

			value, isPresent := newBst.Get(25)
			assert.Equal(t, "jkl", value)
			assert.True(t, isPresent)

			value2, isPresent2 := newBst.Get(10)
			assert.Empty(t, value2)
			assert.True(t, isPresent2)

		})
		t.Run("return empty string and false when value is not found", func(t *testing.T) {
			rootNode := &tree.TreeNode{
				Key:   50,
				Value: "abc",
			}

			newBst := tree.NewBinarySearchTree(rootNode)

			newBst.Set(100, "def")
			newBst.Set(25, "jkl")
			newBst.Set(10, "")

			value, isPresent := newBst.Get(35)
			assert.Empty(t, value)
			assert.False(t, isPresent)
		})
	})
	t.Run("Set", func(t *testing.T) {
		t.Run("insert a node if the key is not present", func(t *testing.T) {
			rootNode := &tree.TreeNode{
				Key:   50,
				Value: "abc",
			}
			newBst := tree.NewBinarySearchTree(rootNode)

			value, isSet := newBst.Set(105, "ddd")
			assert.Empty(t, value)
			assert.True(t, isSet)

			getValue, isPresent := newBst.Get(105)
			assert.Equal(t, "ddd", getValue)
			assert.True(t, isPresent)
		})
		t.Run("if the key is present, insert new value and return old value", func(t *testing.T) {
			rootNode := &tree.TreeNode{
				Key:   50,
				Value: "abc",
			}
			newBst := tree.NewBinarySearchTree(rootNode)

			setValue, _ := newBst.Set(105, "originalValue")
			assert.Empty(t, setValue)

			oldValue, isSet2 := newBst.Set(105, "newValue")
			assert.Equal(t, "originalValue", oldValue)
			assert.True(t, isSet2)

			getValue, isPresent := newBst.Get(105)
			assert.Equal(t, "newValue", getValue)
			assert.True(t, isPresent)
		})
		t.Run("able to insert a node if root node is nil", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree(nil)
			answer, isSet := newBst.Set(200, "hhh")
			assert.Empty(t, answer)
			assert.True(t, isSet)

			value, got := newBst.Get(200)
			assert.Equal(t, "hhh", value)
			assert.True(t, got)
		})

	})
	t.Run("Remove", func(t *testing.T) {
		t.Run("return empty string and false if node to be removed not present", func(t *testing.T) {
			rootNode := &tree.TreeNode{
				Key:   100,
				Value: "100",
			}
			newBst := tree.NewBinarySearchTree(rootNode)
			newBst.Set(75, "75")
			newBst.Set(113, "113")

			removed, wasRemoved := newBst.Remove(20)
			assert.Empty(t, removed)
			assert.False(t, wasRemoved)
		})
		t.Run("removes the element with the given key", func(t *testing.T) {
			t.Run("when node to be removed is a leaf node", func(t *testing.T) {
				rootNode := &tree.TreeNode{
					Key:   100,
					Value: "100",
				}
				newBst := tree.NewBinarySearchTree(rootNode)
				newBst.Set(75, "75")
				newBst.Set(113, "113")

				removed, wasRemoved := newBst.Remove(75)
				assert.Equal(t, "75", removed)
				assert.True(t, wasRemoved)

				getValue, isGot := newBst.Get(75)
				assert.Empty(t, getValue)
				assert.False(t, isGot)
			})
			t.Run("when node to be removed is a leaf node", func(t *testing.T) {

			})
			t.Run("when node to be removed is a leaf node", func(t *testing.T) {

			})
		})
	})
}
