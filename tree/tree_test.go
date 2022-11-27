package tree

import (
	"fmt"
	"strings"
	"testing"
)

func TestTreeAdd(t *testing.T) {
	root := &TreeNode[string]{
		Value: "A",
	}

	b := root.Add("B")
	root.Add("C")
	d := root.Add("D")

	b.Add("E")
	b.Add("F")

	d.Add("G")

	var sb strings.Builder
	root.Preorder(func(val string) {
		sb.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(sb.String())

	sb.Reset()
	root.Postorder(func(val string) {
		sb.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(sb.String())

	sb.Reset()
	root.BFS(func(val string) {
		sb.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(sb.String())

	sb.Reset()
	root.DFS(func(val string) {
		sb.WriteString(fmt.Sprint(val, " - "))
	})
	t.Log(sb.String())
}
