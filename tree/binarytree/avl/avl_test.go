package avl

import (
	"datastructure/tree/drawer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawTest(t *testing.T) {
	tree := &Tree[int, int]{}
	tree.Add(5, 5)
	tree.Add(3, 3)
	tree.Add(2, 2)
	tree.Add(4, 4)
	tree.Add(7, 7)
	tree.Add(6, 6)
	tree.Add(8, 8)

	err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	assert.Nil(t, err)

	value, found := tree.Get(7)
	assert.True(t, found)
	assert.Equal(t, 7, value)
	value, found = tree.Get(6)
	assert.True(t, found)
	assert.Equal(t, 6, value)
	_, found = tree.Get(11)
	assert.False(t, found)
	_, found = tree.Get(1)
	assert.False(t, found)
}

func TestRemoveTest(t *testing.T) {
	tree := &Tree[int, int]{}
	tree.Add(6, 6)
	tree.Add(4, 4)
	tree.Add(3, 3)
	tree.Add(5, 5)

	tree.Remove(4)

	err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	assert.Nil(t, err)

	_, found := tree.Get(6)
	assert.True(t, found)
	_, found = tree.Get(4)
	assert.False(t, found)
	_, found = tree.Get(3)
	assert.True(t, found)
	_, found = tree.Get(5)
	assert.True(t, found)
}

func TestUnbalancedTreeTest(t *testing.T) {
	tree := &Tree[int, int]{}

	for i := 1; i <= 10; i++ {
		tree.Add(i, i)
	}

	err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	assert.Nil(t, err)
}

func TestTreeRemove(t *testing.T) {
	tree := &Tree[int, int]{}

	for i := 2; i <= 8; i++ {
		tree.Add(i, i)
	}

	err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	assert.Nil(t, err)

	removed := tree.Remove(4)
	assert.True(t, removed)

	err = drawer.SaveTreeGraph(tree.Root, "./tree2.png")
	assert.Nil(t, err)

	removed = tree.Remove(3)
	assert.True(t, removed)

	err = drawer.SaveTreeGraph(tree.Root, "./tree3.png")
	assert.Nil(t, err)

	removed = tree.Remove(2)
	assert.True(t, removed)

	err = drawer.SaveTreeGraph(tree.Root, "./tree4.png")
	assert.Nil(t, err)

	removed = tree.Remove(8)
	assert.True(t, removed)

	err = drawer.SaveTreeGraph(tree.Root, "./tree5.png")
	assert.Nil(t, err)
}
