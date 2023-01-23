package tree

import "datastructure/tree/nodeinterface"

type TreeNode[T any] struct {
	Value T

	Childs []*TreeNode[T]
}

func (t *TreeNode[T]) GetChilds() []nodeinterface.Node {
	var childs []nodeinterface.Node
	for _, c := range t.Childs {
		childs = append(childs, c)
	}
	return childs
}

func (t *TreeNode[T]) GetValue() any {
	return t.Value
}

func (t *TreeNode[T]) Add(val T) *TreeNode[T] {
	n := &TreeNode[T]{
		Value: val,
	}

	t.Childs = append(t.Childs, n)
	return n
}

func (t *TreeNode[T]) DFS(fn func(val T)) {
	stack := []*TreeNode[T]{}
	stack = append(stack, t)

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fn(last.Value)

		stack = append(stack, last.Childs...)
	}
}

func (t *TreeNode[T]) Preorder(fn func(val T)) {
	if t == nil {
		return
	}
	fn(t.Value)

	for _, n := range t.Childs {
		n.Preorder(fn)
	}
}

func (t *TreeNode[T]) Postorder(fn func(val T)) {
	if t == nil {
		return
	}

	for _, n := range t.Childs {
		n.Postorder(fn)
	}

	fn(t.Value)
}

func (t *TreeNode[T]) BFS(fn func(val T)) {
	queue := make([]*TreeNode[T], 0)
	queue = append(queue, t)

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		fn(front.Value)

		queue = append(queue, front.Childs...)
	}
}
