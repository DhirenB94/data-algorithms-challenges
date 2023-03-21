package tree_test

import (
	tree "data/challenge-3-map-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		t.Run("return true when the element is found and false when it is not", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree()

			newBst.Set(100, "def")
			newBst.Set(75, "ghi")
			newBst.Set(25, "jkl")
			newBst.Set(1, "mno")

			has := newBst.Has(75)
			assert.True(t, has)
			has = newBst.Has(99)
			assert.False(t, has)
		})
	})
	t.Run("Get", func(t *testing.T) {
		t.Run("return the value and true when the element is found", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree()
			expectedValue := "ghi"

			newBst.Set(100, "def")
			newBst.Set(75, "ghi")
			newBst.Set(25, "jkl")
			newBst.Set(1, "mno")

			getValue, isGot := newBst.Get(75)
			assert.Equal(t, expectedValue, getValue)
			assert.True(t, isGot)
		})
		t.Run("return empty string and false when value is not found", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree()
			expectedValue := ""

			newBst.Set(100, "def")
			newBst.Set(75, "ghi")
			newBst.Set(25, "jkl")
			newBst.Set(1, "mno")

			getValue, isGot := newBst.Get(35)
			assert.Equal(t, expectedValue, getValue)
			assert.False(t, isGot)
		})
	})
	t.Run("Set", func(t *testing.T) {
		t.Run("return an empty string and true when node is not present", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree()
			expectedValue := "def"
			expecteSetValue := ""

			setValue, isSet := newBst.Set(100, "def")
			assert.Equal(t, expecteSetValue, setValue)
			assert.True(t, isSet)

			getValue, isGot := newBst.Get(100)
			assert.Equal(t, expectedValue, getValue)
			assert.True(t, isGot)
		})
		t.Run("if the key is present, insert new value and return old value and true", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree()
			expectedOldValue := "def"
			expectedGetValue := "ghi"

			newBst.Set(100, "def")
			oldValue, isSet := newBst.Set(100, "ghi")
			assert.Equal(t, expectedOldValue, oldValue)
			assert.True(t, isSet)

			getValue, isGot := newBst.Get(100)
			assert.Equal(t, expectedGetValue, getValue)
			assert.True(t, isGot)
		})
	})
	t.Run("Remove", func(t *testing.T) {
		t.Run("return empty string and false if node to be removed not present", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree()

			newBst.Set(100, "def")
			newBst.Set(75, "ghi")
			newBst.Set(25, "jkl")
			newBst.Set(1, "mno")

			removedValue, isRemoved := newBst.Remove(21)
			assert.Empty(t, removedValue)
			assert.False(t, isRemoved)
		})
		t.Run("removes the element with the given key", func(t *testing.T) {
			t.Run("when node to be removed is a leaf node", func(t *testing.T) {
				newBst := tree.NewBinarySearchTree()
				expectedRemovedValue := "75"

				newBst.Set(100, "100")
				newBst.Set(75, "75")
				newBst.Set(113, "113")

				removedValue, isRemoved := newBst.Remove(75)
				assert.Equal(t, expectedRemovedValue, removedValue)
				assert.True(t, isRemoved)

				has := newBst.Has(75)
				assert.False(t, has)
			})
			t.Run("when node to be removed has one child node", func(t *testing.T) {
				newBst := tree.NewBinarySearchTree()
				expectedRemovedValue := "75"

				newBst.Set(100, "100")
				newBst.Set(75, "75")
				newBst.Set(66, "66")
				newBst.Set(113, "113")

				removedValue, isRemoved := newBst.Remove(75)
				assert.Equal(t, expectedRemovedValue, removedValue)
				assert.True(t, isRemoved)

				has := newBst.Has(75)
				assert.False(t, has)

				has = newBst.Has(66)
				assert.True(t, has)
			})
			t.Run("when node to be removed has left and right child nodes", func(t *testing.T) {
				newBst := tree.NewBinarySearchTree()
				expectedRemovedValue := "90"

				newBst.Set(100, "100")
				newBst.Set(75, "75")
				newBst.Set(150, "150")
				newBst.Set(60, "60")
				newBst.Set(90, "90")
				newBst.Set(85, "85")
				newBst.Set(95, "95")
				newBst.Set(92, "92")

				removedValue, isremoved := newBst.Remove(90)
				assert.Equal(t, expectedRemovedValue, removedValue)
				assert.True(t, isremoved)

				getValue, isGot := newBst.Get(90)
				assert.Empty(t, getValue)
				assert.False(t, isGot)

				newBst.Remove(100)

				has := newBst.Has(85)
				assert.True(t, has)

				has = newBst.Has(92)
				assert.True(t, has)

				has = newBst.Has(150)
				assert.True(t, has)
			})
		})
	})
}
