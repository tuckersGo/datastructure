package avl

import (
	"datastructure/tree/nodeinterface"

	"golang.org/x/exp/constraints"
)

type Tree[K constraints.Ordered, V any] struct {
	Root *TreeNode[K, V]
}

type TreeNode[K constraints.Ordered, V any] struct {
	Key   K
	Value V
	depth int

	Left  *TreeNode[K, V]
	Right *TreeNode[K, V]
}

func (n *TreeNode[K, V]) GetChilds() []nodeinterface.Node {
	var childs []nodeinterface.Node
	if n.Left == nil {
		childs = append(childs, nil)
	} else {
		childs = append(childs, n.Left)
	}

	if n.Right == nil {
		childs = append(childs, nil)
	} else {
		childs = append(childs, n.Right)
	}
	return childs
}

func (n *TreeNode[K, V]) GetValue() any {
	return n.Value
}

func newNode[K constraints.Ordered, V any](key K, value V) *TreeNode[K, V] {
	return &TreeNode[K, V]{
		Key:   key,
		Value: value,
		depth: 1,
	}
}

func (t *Tree[K, V]) Add(key K, value V) {
	t.Root = add(t.Root, key, value)
}

func add[K constraints.Ordered, V any](n *TreeNode[K, V], key K, value V) *TreeNode[K, V] {
	if n == nil {
		return newNode(key, value)
	}

	if key < n.Key {
		n.Left = add(n.Left, key, value)
	} else {
		n.Right = add(n.Right, key, value)
	}
	return rotateAfterInsert(n, key)
}

func (n *TreeNode[K, V]) getBalance() int {
	if n == nil {
		return 0
	}
	return n.Left.getDepth() - n.Right.getDepth()
}

func (n *TreeNode[K, V]) getDepth() int {
	if n == nil {
		return 0
	}
	return n.depth
}

func rotateAfterInsert[K constraints.Ordered, V any](n *TreeNode[K, V], key K) *TreeNode[K, V] {
	n.updateDepth()

	balance := n.getBalance()

	if balance >= -1 && balance <= 1 {
		return n
	}

	if balance > 1 && key < n.Left.Key {
		// type 1 - LL
		return rightRotate(n)
	}
	if balance < -1 && key >= n.Right.Key {
		// type 1 - RR
		return leftRotate(n)
	}
	if balance > 1 && key >= n.Left.Key {
		// type 2 - LR
		n.Left = leftRotate(n.Left)
		return rightRotate(n)
	}
	if balance < -1 && key < n.Right.Key {
		// type 2 - RL
		n.Right = rightRotate(n.Right)
		return leftRotate(n)
	}
	return n
}

func rightRotate[K constraints.Ordered, V any](n *TreeNode[K, V]) *TreeNode[K, V] {
	l := n.Left
	t := l.Right

	l.Right = n
	n.Left = t

	n.updateDepth()
	l.updateDepth()
	return l
}

func leftRotate[K constraints.Ordered, V any](n *TreeNode[K, V]) *TreeNode[K, V] {
	r := n.Right
	t := r.Left

	r.Left = n
	n.Right = t

	n.updateDepth()
	r.updateDepth()
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (n *TreeNode[K, V]) updateDepth() {
	n.depth = max(n.Left.getDepth(), n.Right.getDepth()) + 1
}

func (t *Tree[K, V]) Get(key K) (value V, found bool) {
	if t.Root == nil {
		found = false
		return
	}

	return t.Root.get(key)
}

func (n *TreeNode[K, V]) get(key K) (value V, found bool) {
	if n == nil {
		found = false
		return
	}
	if n.Key == key {
		value = n.Value
		found = true
		return
	}
	if key < n.Key {
		return n.Left.get(key)
	} else {
		return n.Right.get(key)
	}
}

func (t *Tree[K, V]) Remove(key K) bool {
	if t.Root == nil {
		return false
	}

	var removed bool
	t.Root, removed = remove(t.Root, key)
	return removed
}

func remove[K constraints.Ordered, V any](n *TreeNode[K, V], key K) (*TreeNode[K, V], bool) {
	if n == nil {
		return nil, false
	}

	var removed bool
	if key == n.Key {
		removed = true
		if n.Left == nil && n.Right == nil {
			n = nil
		} else if n.Left == nil {
			n = n.Right
		} else if n.Right == nil {
			n = n.Left
		} else {
			leftMostNode := n.findAndRemoveLeftMostNode()
			leftMostNode.Right = n.Right
			n = leftMostNode
		}
	} else {
		if key < n.Key {
			n.Left, removed = remove(n.Left, key)
		} else {
			n.Right, removed = remove(n.Right, key)
		}
	}

	if removed && n != nil {
		n = rotateDelete(n)
	}
	return n, removed
}

func rotateDelete[K constraints.Ordered, V any](n *TreeNode[K, V]) *TreeNode[K, V] {
	n.updateDepth()
	balance := n.getBalance()

	if balance > 1 && n.Left.getBalance() >= 0 {
		// LL - type1
		return rightRotate(n)
	}
	if balance > 1 && n.Left.getBalance() < 0 {
		// LR - type2
		n.Left = leftRotate(n.Left)
		return rightRotate(n)
	}
	if balance < -1 && n.Right.getBalance() <= 0 {
		// RR - type1
		return leftRotate(n)
	}
	if balance < -1 && n.Right.getBalance() > 0 {
		// RL - type2
		n.Right = rightRotate(n.Right)
		return leftRotate(n)
	}
	return n
}

func (n *TreeNode[K, V]) remove(key K) (*TreeNode[K, V], bool) {
	if n.Key == key {
		if n.Left == nil && n.Right == nil {
			return nil, true
		}

		if n.Left == nil {
			return n.Right, true
		}
		if n.Right == nil {
			return n.Left, true
		}

		// Find left-most node
		leftMostNode := n.findAndRemoveLeftMostNode()
		leftMostNode.Right = n.Right
		return leftMostNode, true
	}

	var removed bool
	if key < n.Key {
		if n.Left == nil {
			return n, false
		}
		n.Left, removed = n.Left.remove(key)
		return n, removed
	} else {
		if n.Right == nil {
			return n, false
		}
		n.Right, removed = n.Right.remove(key)
		return n, removed
	}
}

func (n *TreeNode[K, V]) findAndRemoveLeftMostNode() *TreeNode[K, V] {
	if n.Left == nil {
		return nil
	}

	rightMost := n.Left
	parent := n

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
