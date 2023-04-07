package tree_test

import (
	interfaces "data"
	tree "data/challenge-3-map-tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		t.Run("return true when the element is found and false when it is not", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree[interfaces.MyCustomInt, interfaces.MyCustomString]()

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
		t.Run("return the value the element is found", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree[interfaces.MyCustomInt, interfaces.MyCustomString]()
			expectedValue := "ghi"

			newBst.Set(100, "def")
			newBst.Set(75, "ghi")
			newBst.Set(25, "jkl")
			newBst.Set(1, "mno")

			getValue := newBst.Get(75)
			assert.Equal(t, expectedValue, string(*getValue))
		})
		t.Run("return nil when value is not found", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree[interfaces.MyCustomInt, interfaces.MyCustomString]()

			newBst.Set(100, "def")
			newBst.Set(75, "ghi")
			newBst.Set(25, "jkl")
			newBst.Set(1, "mno")

			getValue := newBst.Get(35)
			assert.Nil(t, getValue)
		})
	})
	t.Run("Set", func(t *testing.T) {
		t.Run("when a node is not already present, return nil and set the new node", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree[interfaces.MyCustomInt, interfaces.MyCustomString]()
			expectedValue := "def"

			setValue := newBst.Set(100, "def")
			assert.Nil(t, setValue)

			getValue := newBst.Get(100)
			assert.Equal(t, expectedValue, string(*getValue))
		})
		t.Run("if node with key already exists, update its value and return the old value", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree[interfaces.MyCustomInt, interfaces.MyCustomString]()
			expectedOldValue := "def"
			expectedGetValue := "ghi"

			newBst.Set(100, "def")
			oldValue := newBst.Set(100, "ghi")
			assert.Equal(t, expectedOldValue, string(*oldValue))

			getValue := newBst.Get(100)
			assert.Equal(t, expectedGetValue, string(*getValue))
		})
	})
	t.Run("Remove", func(t *testing.T) {
		t.Run("return nil if node to be removed not present", func(t *testing.T) {
			newBst := tree.NewBinarySearchTree[interfaces.MyCustomInt, interfaces.MyCustomString]()

			newBst.Set(100, "def")
			newBst.Set(75, "ghi")
			newBst.Set(25, "jkl")
			newBst.Set(1, "mno")

			removedValue := newBst.Remove(21)
			assert.Nil(t, removedValue)
		})
		//t.Run("removes the element with the given key", func(t *testing.T) {
		//t.Run("when node to be removed is a leaf node", func(t *testing.T) {
		//	newBst := tree.NewBinarySearchTree[interfaces.MyCustomInt, interfaces.MyCustomString]()
		//	expectedRemovedValue := "75"
		//
		//	newBst.Set(100, "100")
		//	newBst.Set(75, "75")
		//	newBst.Set(113, "113")
		//
		//	removedValue := newBst.Remove(75)
		//	assert.Equal(t, expectedRemovedValue, string(*removedValue))
		//
		//	has := newBst.Has(75)
		//	assert.False(t, has)
		//})
		//t.Run("when node to be removed has one child node", func(t *testing.T) {
		//	newBst := tree.NewBinarySearchTree[string]()
		//	expectedRemovedValue := "75"
		//
		//	newBst.Set(100, "100")
		//	newBst.Set(75, "75")
		//	newBst.Set(66, "66")
		//	newBst.Set(113, "113")
		//
		//	removedValue := newBst.Remove(75)
		//	assert.Equal(t, expectedRemovedValue, *removedValue)
		//
		//	has := newBst.Has(75)
		//	assert.False(t, has)
		//
		//	has = newBst.Has(66)
		//	assert.True(t, has)
		//})
		//		t.Run("when node to be removed has left and right child nodes", func(t *testing.T) {
		//			newBst := tree.NewBinarySearchTree[string]()
		//			expectedRemovedValue := "90"
		//
		//			newBst.Set(100, "100")
		//			newBst.Set(75, "75")
		//			newBst.Set(150, "150")
		//			newBst.Set(60, "60")
		//			newBst.Set(90, "90")
		//			newBst.Set(85, "85")
		//			newBst.Set(95, "95")
		//			newBst.Set(92, "92")
		//
		//			gotVal := newBst.Get(90)
		//			fmt.Println("GET VALUE", *gotVal)
		//			removedValue := newBst.Remove(90)
		//			fmt.Println("REMOVED VAL ", *removedValue)
		//			assert.Equal(t, expectedRemovedValue, *removedValue)
		//
		//			getValue := newBst.Get(90)
		//			assert.Nil(t, getValue)
		//
		//			newBst.Remove(100)
		//
		//			has := newBst.Has(85)
		//			assert.True(t, has)
		//
		//			has = newBst.Has(92)
		//			assert.True(t, has)
		//
		//			has = newBst.Has(150)
		//			assert.True(t, has)
		//		})
		//})
	})
}
