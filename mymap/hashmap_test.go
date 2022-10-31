package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap(t *testing.T) {
	var h HashMap[int]
	h.Add("tucker", 100)

	val, ok := h.Get("tucker")
	assert.True(t, ok)
	assert.Equal(t, 100, val)

	h.Add("golang", 200)

	val, ok = h.Get("golang")
	assert.True(t, ok)
	assert.Equal(t, 200, val)

	val, ok = h.Get("tucker")
	assert.True(t, ok)
	assert.Equal(t, 100, val)

	h.Add("awesome", 300)

	val, ok = h.Get("awesome")
	assert.True(t, ok)
	assert.Equal(t, 300, val)

	val, ok = h.Get("golang")
	assert.True(t, ok)
	assert.Equal(t, 200, val)

	val, ok = h.Get("tucker")
	assert.True(t, ok)
	assert.Equal(t, 100, val)

}

func TestGoBasicMap(t *testing.T) {

	m := make(map[string]int)
	m["tucker"] = 100
	m["golang"] = 200
	m["awesome"] = 300
	m["bbb"] = 0

	assert.Equal(t, 100, m["tucker"])
	assert.Equal(t, 200, m["golang"])
	assert.Equal(t, 300, m["awesome"])
	assert.Equal(t, 0, m["aaa"])

	_, ok := m["aaa"]
	assert.False(t, ok)

	delete(m, "tucker")
	v, ok := m["tucker"]
	assert.Equal(t, 0, v)
	assert.False(t, ok)

}
