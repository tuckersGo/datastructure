package mymap

import "hash/crc32"

const arraySize = 3571

type hashData[T any] struct {
	key   string
	value T
}

type HashMap[T any] struct {
	arr [arraySize][]hashData[T]
}

func hashfn(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (h *HashMap[T]) Add(key string, value T) {
	hash := hashfn(key)
	var hd = hashData[T]{
		key:   key,
		value: value,
	}
	h.arr[hash%arraySize] = append(h.arr[hash%arraySize], hd)
}

func (h *HashMap[T]) Get(key string) (T, bool) {
	hash := hashfn(key)
	for _, hd := range h.arr[hash%arraySize] {
		if hd.key == key {
			return hd.value, true
		}
	}
	var tempT T
	return tempT, false
}
