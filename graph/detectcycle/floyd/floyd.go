package floyd

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func (n *Node[T]) Link(v T) *Node[T] {
	if n == nil {
		return nil
	}
	n.Next = &Node[T]{
		Value: v,
	}
	return n.Next
}

// if there is cycle then return true
func DetectLoop[T any](head *Node[T]) bool {
	var slow, fast *Node[T]
	slow = head
	fast = head

	for slow != nil &&
		fast != nil &&
		fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
