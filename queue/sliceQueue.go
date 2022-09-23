package queue

type SliceQueue[T any] struct {
	arr []T
}

func NewSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{}
}

func (q *SliceQueue[T]) Push(val T) {
	q.arr = append(q.arr, val)
}

func (q *SliceQueue[T]) Pop() T {
	var front T
	if len(q.arr) == 0 {
		return front
	}
	front = q.arr[0]
	q.arr = q.arr[1:]
	return front
}
