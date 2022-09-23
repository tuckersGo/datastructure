package linkedlist

import (
	"datastructure/linkedlist/doublelinkedlist"
	"datastructure/linkedlist/singlelinkedlist"
	"testing"
)

func BenchmarkSingleLinkedList(b *testing.B) {
	var l singlelinkedlist.LinkedList[int]
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse()
}

func BenchmarkSingleLinkedListSwap(b *testing.B) {
	var l singlelinkedlist.LinkedList[int]
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse2()
}

func BenchmarkDoubleLinkedList(b *testing.B) {
	var l doublelinkedlist.LinkedList[int]
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse()
}
