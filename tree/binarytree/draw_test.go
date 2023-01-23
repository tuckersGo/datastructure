package binarytree

import (
	"datastructure/tree/drawer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawTest(t *testing.T) {
	root := &TreeNode{
		Value: "A",
	}

	b := &TreeNode{
		Value: "B",
	}
	root.Left = b

	c := &TreeNode{
		Value: "C",
	}
	root.Right = c

	d := &TreeNode{
		Value: "D",
	}
	b.Left = d

	e := &TreeNode{
		Value: "E",
	}
	b.Right = e

	f := &TreeNode{
		Value: "F",
	}
	c.Left = f

	g := &TreeNode{
		Value: "G",
	}
	c.Right = g

	h := &TreeNode{
		Value: "H",
	}
	d.Left = h

	i := &TreeNode{
		Value: "I",
	}
	g.Right = i

	err := drawer.SaveTreeGraph(root, "./tree.png")
	assert.Nil(t, err)
}
