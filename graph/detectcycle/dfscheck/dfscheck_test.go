package dfscheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoop(t *testing.T) {
	root := &Node[string]{
		Value: "A",
	}
	b := root.Link("B")
	c := root.Link("C")
	d := root.Link("D")
	d.LinkNode(c)
	b.Link("E")
	f := c.Link("F")
	f.LinkNode(d)

	assert.True(t, DetectLoop(root))
}

func TestNotLoo(t *testing.T) {
	root := &Node[string]{
		Value: "A",
	}

	b := root.Link("B")
	c := root.Link("C")
	d := root.Link("D")
	b.Link("E")
	f := c.Link("F")
	f.LinkNode(d)

	assert.False(t, DetectLoop(root))
}
