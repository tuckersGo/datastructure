package stack

import (
	"container/list"
)

type ListStack[T any] struct {
	l *list.List
}

func NewList[T any]() *ListStack[T] {
	return &ListStack[T]{
		l: list.New(),
	}
}

func (q *ListStack[T]) Push(val T) {
	q.l.PushBack(val)
}

func (q *ListStack[T]) Pop() T {
	front := q.l.Back()
	q.l.Remove(front)
	return front.Value.(T)
}
