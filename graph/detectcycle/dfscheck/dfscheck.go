package dfscheck

type Node[T any] struct {
	Value T
	Links []*Node[T]
}

func (n *Node[T]) LinkNode(l *Node[T]) {
	n.Links = append(n.Links, l)
}

func (n *Node[T]) Link(v T) *Node[T] {
	newNode := &Node[T]{
		Value: v,
	}
	n.Links = append(n.Links, newNode)
	return newNode
}

func DetectLoop[T any](head *Node[T]) bool {
	ancestors := make(map[*Node[T]]bool)
	ancestors[head] = true

	return internalDetectLoop(head, ancestors)
}

func internalDetectLoop[T any](node *Node[T], ancestors map[*Node[T]]bool) bool {
	ancestors[node] = true
	for _, l := range node.Links {
		if ancestors[l] {
			return true
		}

		if internalDetectLoop(l, ancestors) {
			return true
		}
	}
	delete(ancestors, node)
	return false
}
