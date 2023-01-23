package bst

import (
	"datastructure/tree/drawer"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LesserInt int

func (a LesserInt) Less(b Lesser) bool {
	bvalue, ok := b.(LesserInt)
	if !ok {
		panic(fmt.Sprintf("b should be LesserInt type. type:%T", b))
	}

	return a < bvalue
}

func TestDrawTest(t *testing.T) {
	tree := &Tree{}
	tree.Add(LesserInt(5))
	tree.Add(LesserInt(3))
	tree.Add(LesserInt(2))
	tree.Add(LesserInt(4))
	tree.Add(LesserInt(7))
	tree.Add(LesserInt(6))
	tree.Add(LesserInt(8))

	err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	assert.Nil(t, err)

	assert.True(t, tree.Contains(LesserInt(7)))
	assert.True(t, tree.Contains(LesserInt(6)))
	assert.True(t, tree.Contains(LesserInt(2)))
	assert.False(t, tree.Contains(LesserInt(11)))
	assert.False(t, tree.Contains(LesserInt(1)))
}

func TestRemoveTest(t *testing.T) {
	tree := &Tree{}
	tree.Add(LesserInt(6))
	tree.Add(LesserInt(4))
	tree.Add(LesserInt(3))
	tree.Add(LesserInt(5))

	tree.Remove(LesserInt(4))

	err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	assert.Nil(t, err)

	assert.False(t, tree.Contains(LesserInt(4)))
	assert.True(t, tree.Contains(LesserInt(6)))
	assert.True(t, tree.Contains(LesserInt(3)))
	assert.True(t, tree.Contains(LesserInt(5)))
}

func TestRemoveTest2(t *testing.T) {
	tree := &Tree{}
	tree.Add(LesserInt(6))
	tree.Add(LesserInt(3))
	tree.Add(LesserInt(2))

	tree.Remove(LesserInt(3))

	err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	assert.Nil(t, err)

	assert.False(t, tree.Contains(LesserInt(3)))
	assert.True(t, tree.Contains(LesserInt(6)))
	assert.True(t, tree.Contains(LesserInt(2)))
}

func TestRemoveTest3(t *testing.T) {
	tree := &Tree{}
	tree.Add(LesserInt(2))
	tree.Add(LesserInt(3))
	tree.Add(LesserInt(4))

	tree.Remove(LesserInt(3))

	err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	assert.Nil(t, err)

	assert.False(t, tree.Contains(LesserInt(3)))
	assert.True(t, tree.Contains(LesserInt(2)))
	assert.True(t, tree.Contains(LesserInt(4)))
}

func TestUnbalancedTreeTest(t *testing.T) {
	tree := &Tree{}
	for i := 10; i >= 1; i-- {
		tree.Add(LesserInt(i))
	}

	err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	assert.Nil(t, err)
}
