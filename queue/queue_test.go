package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 3, s.Pop())
}

func TestPush2(t *testing.T) {
	s := NewSliceQueue[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 3, s.Pop())
}

func TestPush3(t *testing.T) {
	s := NewList[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 3, s.Pop())
}

func BenchmarkLinkedListQueue(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkSliceQueue(b *testing.B) {
	s := NewSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkListQueue(b *testing.B) {
	s := NewList[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
