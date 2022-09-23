package queue

import "datastructure/linkedlist/doublelinkedlist"

type Queue[T any] struct {
	l *doublelinkedlist.LinkedList[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		l: &doublelinkedlist.LinkedList[T]{},
	}
}

func (q *Queue[T]) Push(val T) {
	q.l.PushBack(val)
}

func (q *Queue[T]) Pop() T {
	return q.l.PopFront().Value
}
