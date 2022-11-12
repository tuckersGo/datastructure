package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedMap(t *testing.T) {
	var s SortedMap[string, int]

	s.Add("bbb", 10)

	v, ok := s.Get("bbb")
	assert.True(t, ok)
	assert.Equal(t, 10, v)

	s.Add("aaa", 20)

	v, ok = s.Get("aaa")
	assert.True(t, ok)
	assert.Equal(t, 20, v)

	assert.Equal(t, "aaa", s.Arr[0].Key)
	assert.Equal(t, "bbb", s.Arr[1].Key)
}

func TestSortedMapOverlapped(t *testing.T) {
	var s SortedMap[string, int]

	s.Add("bbb", 10)

	v, ok := s.Get("bbb")
	assert.True(t, ok)
	assert.Equal(t, 10, v)

	s.Add("bbb", 20)

	v, ok = s.Get("bbb")
	assert.True(t, ok)
	assert.Equal(t, 20, v)

	assert.Equal(t, 1, len(s.Arr))

}

func TestSortedGetEmpty(t *testing.T) {
	var s SortedMap[string, int]

	v, ok := s.Get("aaa")
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestSortedMapRemove(t *testing.T) {
	var s SortedMap[string, int]

	s.Add("cccc", 30)
	s.Add("bbb", 20)
	s.Add("aaa", 10)

	removed := s.Remove("bbb")
	assert.True(t, removed)
	removed = s.Remove("bbb")
	assert.False(t, removed)

	assert.Equal(t, 2, len(s.Arr))

	assert.Equal(t, "aaa", s.Arr[0].Key)
	assert.Equal(t, "cccc", s.Arr[1].Key)
}
