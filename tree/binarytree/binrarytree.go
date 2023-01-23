package binarytree

import "datastructure/tree/nodeinterface"

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value any

	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) GetChilds() []nodeinterface.Node {
	var childs []nodeinterface.Node
	if t.Left == nil {
		childs = append(childs, nil)
	} else {
		childs = append(childs, t.Left)
	}

	if t.Right == nil {
		childs = append(childs, nil)
	} else {
		childs = append(childs, t.Right)
	}
	return childs
}

func (t *TreeNode) GetValue() any {
	return t.Value
}
