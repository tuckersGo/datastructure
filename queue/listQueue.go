package queue

import (
	"container/list"
)

type ListQueue[T any] struct {
	l *list.List
}

func NewList[T any]() *ListQueue[T] {
	return &ListQueue[T]{
		l: list.New(),
	}
}

func (q *ListQueue[T]) Push(val T) {
	q.l.PushBack(val)
}

func (q *ListQueue[T]) Pop() T {
	front := q.l.Front()
	q.l.Remove(front)
	return front.Value.(T)
}
