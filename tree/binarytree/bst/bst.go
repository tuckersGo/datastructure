package bst

import (
	"datastructure/tree/nodeinterface"
)

type Lesser interface {
	Less(other Lesser) bool
}

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value Lesser

	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) GetChilds() []nodeinterface.Node {
	var childs []nodeinterface.Node
	if t.Left == nil {
		childs = append(childs, nil)
	} else {
		childs = append(childs, t.Left)
	}

	if t.Right == nil {
		childs = append(childs, nil)
	} else {
		childs = append(childs, t.Right)
	}
	return childs
}

func (t *TreeNode) GetValue() any {
	return t.Value
}

func (t *Tree) Add(value Lesser) *TreeNode {
	if t.Root == nil {
		t.Root = &TreeNode{
			Value: value,
		}
		return t.Root
	}

	return t.Root.add(value)
}

func (t *TreeNode) add(value Lesser) *TreeNode {
	if t.Value.Less(value) {
		if t.Right == nil {
			t.Right = &TreeNode{
				Value: value,
			}
			return t.Right
		}
		return t.Right.add(value)
	} else {
		if t.Left == nil {
			t.Left = &TreeNode{
				Value: value,
			}
			return t.Left
		}
		return t.Left.add(value)
	}
}

func (t *Tree) Contains(value Lesser) bool {
	if t.Root == nil {
		return false
	}

	return t.Root.contains(value)
}

func (t *TreeNode) contains(value Lesser) bool {
	if equal(t.Value, value) {
		return true
	}
	if t.Value.Less(value) {
		// my < value
		if t.Right == nil {
			return false
		}
		return t.Right.contains(value)
	} else {
		// my >= value
		if value.Less(t.Value) {
			// value < my
			// my > value
			if t.Left == nil {
				return false
			}
			return t.Left.contains(value)
		} else {
			// my == value
			return true
		}
	}
}

func equal(a, b Lesser) bool {
	if a.Less(b) {
		// a < b
		return false
	}
	if b.Less(a) {
		// b < a
		return false
	}
	return true
}

func (t *Tree) Remove(value Lesser) bool {
	if t.Root == nil {
		return false
	}

	var removed bool
	t.Root, removed = t.Root.remove(value)
	return removed
}

func (t *TreeNode) remove(value Lesser) (*TreeNode, bool) {
	if equal(t.Value, value) {
		if t.Left == nil && t.Right == nil {
			return nil, true
		}

		if t.Left == nil {
			return t.Right, true
		}
		if t.Right == nil {
			return t.Left, true
		}

		// Find left-most node
		leftMostNode := t.findAndRemoveLeftMostNode()
		leftMostNode.Right = t.Right
		return leftMostNode, true
	}

	var removed bool
	if t.Value.Less(value) {
		// value > t.Value
		if t.Right == nil {
			return t, false
		}
		t.Right, removed = t.Right.remove(value)
		return t, removed
	} else {
		// value < t.Value
		if t.Left == nil {
			return t, false
		}
		t.Left, removed = t.Left.remove(value)
		return t, removed
	}
}

func (t *TreeNode) findAndRemoveLeftMostNode() *TreeNode {
	if t.Left == nil {
		return nil
	}

	rightMost := t.Left
	parent := t

	for rightMost.Right != nil {
		parent = rightMost
		rightMost = rightMost.Right
	}

	if parent.Left == rightMost {
		parent.Left = nil
	} else {
		parent.Right = nil
	}
	return rightMost
}
