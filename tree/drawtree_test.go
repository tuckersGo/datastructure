package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawTest(t *testing.T) {
	root := &TreeNode[string]{
		Value: "A",
	}

	b := root.Add("B")
	root.Add("C")
	d := root.Add("D")

	b.Add("E")
	b.Add("F")

	d.Add("G")

	err := SaveTreeGraph[string](root, "./tree.png")
	assert.Nil(t, err)
}
