package trie

import "datastructure/tree/nodeinterface"

type Node struct {
	Children map[rune]*Node
	Value    string
}

func NewNode(value string) *Node {
	return &Node{
		Children: make(map[rune]*Node),
		Value:    value,
	}
}

func (n *Node) GetChilds() []nodeinterface.Node {
	rst := make([]nodeinterface.Node, len(n.Children))
	idx := 0
	for _, c := range n.Children {
		rst[idx] = c
		idx++
	}
	return rst
}

func (n *Node) GetValue() any {
	return n.Value
}

func Insert(root *Node, key string) bool {
	runes := []rune(key)
	cur := root
	for i, c := range runes {
		node := cur.Children[c]
		if node == nil {
			node = NewNode(string(runes[:i+1]))
			cur.Children[c] = node
		}
		cur = node
	}
	return true
}

func AutoComplete(root *Node, key string) string {
	cur := root
	for _, c := range key {
		if n, ok := cur.Children[c]; ok {
			cur = n
		} else {
			return ""
		}
	}
	for {
		if len(cur.Children) == 0 {
			break
		}
		for _, n := range cur.Children {
			cur = n
			break
		}
	}
	return cur.Value
}
