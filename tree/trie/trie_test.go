package trie

import (
	"datastructure/tree/drawer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	root := NewNode("")
	success := Insert(root, "tea")
	assert.True(t, success)

	success = Insert(root, "ted")
	assert.True(t, success)

	success = Insert(root, "ten")
	assert.True(t, success)

	success = Insert(root, "app")
	assert.True(t, success)

	success = Insert(root, "in")
	assert.True(t, success)

	success = Insert(root, "inn")
	assert.True(t, success)

	err := drawer.SaveTreeGraph(root, "./tree.png")
	assert.Nil(t, err)
}

func TestAutoComplete(t *testing.T) {
	root := NewNode("")
	success := Insert(root, "한국인")
	assert.True(t, success)

	success = Insert(root, "한국 사랑해요")
	assert.True(t, success)

	success = Insert(root, "안녕하세요")
	assert.True(t, success)

	success = Insert(root, "안사요")
	assert.True(t, success)

	res := AutoComplete(root, "한국 사")
	assert.Equal(t, "한국 사랑해요", res)

	res = AutoComplete(root, "안사")
	assert.Equal(t, "안사요", res)

	success = Insert(root, "hello")
	assert.True(t, success)

	success = Insert(root, "hell")
	assert.True(t, success)

	success = Insert(root, "heal")
	assert.True(t, success)

	res = AutoComplete(root, "hea")
	assert.Equal(t, "heal", res)

}
