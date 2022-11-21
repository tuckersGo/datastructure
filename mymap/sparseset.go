package mymap

import "golang.org/x/exp/constraints"

type SparseSet[TKey constraints.Ordered, TValue any] struct {
	dense  []Element[TKey, TValue]
	sparse map[TKey]int
}

func NewSparseSet[TKey constraints.Ordered, TValue any]() *SparseSet[TKey, TValue] {
	return &SparseSet[TKey, TValue]{
		sparse: make(map[TKey]int),
	}
}

func (s *SparseSet[TKey, TValue]) Add(key TKey, value TValue) {
	if idx, ok := s.sparse[key]; ok {
		s.dense[idx].Value = value
		return
	}

	s.dense = append(s.dense, Element[TKey, TValue]{
		Key:   key,
		Value: value,
	})

	s.sparse[key] = len(s.dense) - 1
}

func (s *SparseSet[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	if idx, ok := s.sparse[key]; ok {
		value = s.dense[idx].Value
		found = true
		return
	}

	found = false
	return
}

func (s *SparseSet[TKey, TValue]) Remove(key TKey) bool {
	if idx, ok := s.sparse[key]; ok {
		last := len(s.dense) - 1
		if idx < last {
			s.dense[idx] = s.dense[last]
			s.sparse[s.dense[last].Key] = idx
		}
		s.dense = s.dense[:last]
		delete(s.sparse, key)
		return true
	}

	return false
}

type Iterator[TKey constraints.Ordered, TValue any] struct {
	dense []Element[TKey, TValue]
	idx   int
}

func (s *SparseSet[TKey, TValue]) Iterator() *Iterator[TKey, TValue] {
	return &Iterator[TKey, TValue]{
		dense: s.dense,
		idx:   0,
	}
}

func (i *Iterator[TKey, TValue]) IsEnd() bool {
	return i.idx >= len(i.dense)
}

func (i *Iterator[TKey, TValue]) Next() {
	i.idx++
}

func (i *Iterator[TKey, TValue]) Get() Element[TKey, TValue] {
	return i.dense[i.idx]
}
