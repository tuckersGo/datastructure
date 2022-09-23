package stack

import "datastructure/linkedlist/doublelinkedlist"

type Stack[T any] struct {
	l *doublelinkedlist.LinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		l: &doublelinkedlist.LinkedList[T]{},
	}
}

func (s *Stack[T]) Push(val T) {
	s.l.PushBack(val)
}

func (s *Stack[T]) Pop() T {
	return s.l.PopBack().Value
}
