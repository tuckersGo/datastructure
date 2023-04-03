package calculator

import (
	"datastructure/tree/binarytree"
	"strconv"
)

type token interface {
	String() string
	Evaluate(tokens *[]token) (int, bool)
	Priority() int
	MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool)
	EvaluateTree(left, right *binarytree.TreeNode) int
}

type number int

func (n number) String() string {
	return strconv.Itoa(int(n))
}

func (n number) Evaluate(tokens *[]token) (int, bool) {
	return int(n), true
}

func (n number) Priority() int {
	return 0
}

func (n number) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return &binarytree.TreeNode{
		Value: n,
	}, true
}

func (n number) EvaluateTree(left, right *binarytree.TreeNode) int {
	return int(n)
}

type plus struct{}

func (p plus) String() string {
	return "+"
}

func (p plus) Evaluate(tokens *[]token) (int, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return 0, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	rh, success := top.Evaluate(&newtokens)
	if !success || len(newtokens) == 0 {
		return 0, false
	}
	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	lh, success := top.Evaluate(&newtokens)
	if !success {
		return 0, false
	}

	*tokens = newtokens
	return lh + rh, true
}

func (n plus) Priority() int {
	return 1
}

func makeOpTreeNode(opToken token, tokens *[]token) (*binarytree.TreeNode, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return nil, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	right, success := top.MakeTreeNode(&newtokens)
	if !success || len(newtokens) == 0 {
		return nil, false
	}
	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	left, success := top.MakeTreeNode(&newtokens)
	if !success {
		return nil, false
	}

	*tokens = newtokens
	return &binarytree.TreeNode{
		Value: opToken,
		Left:  left,
		Right: right,
	}, true
}

func (n plus) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return makeOpTreeNode(n, tokens)
}

func (n plus) EvaluateTree(left, right *binarytree.TreeNode) int {
	leftToken := left.Value.(token)
	if leftToken == nil {
		panic("leftToken should be not nil")
	}
	lh := leftToken.EvaluateTree(left.Left, left.Right)

	rightToken := right.Value.(token)
	if rightToken == nil {
		panic("rightToken should be not nil")
	}
	rh := rightToken.EvaluateTree(right.Left, right.Right)

	return lh + rh
}

type minus struct{}

func (m minus) String() string {
	return "-"
}

func (p minus) Evaluate(tokens *[]token) (int, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return 0, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	rh, success := top.Evaluate(&newtokens)
	if !success || len(newtokens) == 0 {
		return 0, false
	}
	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	lh, success := top.Evaluate(&newtokens)
	if !success {
		return 0, false
	}

	*tokens = newtokens
	return lh - rh, true
}

func (n minus) Priority() int {
	return 1
}

func (n minus) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return makeOpTreeNode(n, tokens)
}

func (n minus) EvaluateTree(left, right *binarytree.TreeNode) int {
	leftToken := left.Value.(token)
	if leftToken == nil {
		panic("leftToken should be not nil")
	}
	lh := leftToken.EvaluateTree(left.Left, left.Right)

	rightToken := right.Value.(token)
	if rightToken == nil {
		panic("rightToken should be not nil")
	}
	rh := rightToken.EvaluateTree(right.Left, right.Right)

	return lh - rh
}

type multiple struct{}

func (m multiple) String() string {
	return "*"
}

func (p multiple) Evaluate(tokens *[]token) (int, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return 0, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	rh, success := top.Evaluate(&newtokens)
	if !success || len(newtokens) == 0 {
		return 0, false
	}
	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	lh, success := top.Evaluate(&newtokens)
	if !success {
		return 0, false
	}

	*tokens = newtokens
	return lh * rh, true
}

func (n multiple) Priority() int {
	return 2
}

func (n multiple) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return makeOpTreeNode(n, tokens)
}

func (n multiple) EvaluateTree(left, right *binarytree.TreeNode) int {
	leftToken := left.Value.(token)
	if leftToken == nil {
		panic("leftToken should be not nil")
	}
	lh := leftToken.EvaluateTree(left.Left, left.Right)

	rightToken := right.Value.(token)
	if rightToken == nil {
		panic("rightToken should be not nil")
	}
	rh := rightToken.EvaluateTree(right.Left, right.Right)

	return lh * rh
}

type divide struct{}

func (m divide) String() string {
	return "/"
}

func (p divide) Evaluate(tokens *[]token) (int, bool) {
	newtokens := *tokens
	if len(newtokens) < 2 {
		return 0, false
	}

	top, newtokens := newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	rh, success := top.Evaluate(&newtokens)
	if !success || len(newtokens) == 0 {
		return 0, false
	}
	top, newtokens = newtokens[len(newtokens)-1], newtokens[:len(newtokens)-1]
	lh, success := top.Evaluate(&newtokens)
	if !success {
		return 0, false
	}

	*tokens = newtokens
	return lh / rh, true
}

func (n divide) Priority() int {
	return 2
}

func (n divide) MakeTreeNode(tokens *[]token) (*binarytree.TreeNode, bool) {
	return makeOpTreeNode(n, tokens)
}

func (n divide) EvaluateTree(left, right *binarytree.TreeNode) int {
	leftToken := left.Value.(token)
	if leftToken == nil {
		panic("leftToken should be not nil")
	}
	lh := leftToken.EvaluateTree(left.Left, left.Right)

	rightToken := right.Value.(token)
	if rightToken == nil {
		panic("rightToken should be not nil")
	}
	rh := rightToken.EvaluateTree(right.Left, right.Right)

	return lh / rh
}

type parser struct {
	eval        []rune
	idx         int
	parsedToken token
}

func (p *parser) parse() bool {
	// ignore spaces
	for {
		if p.idx >= len(p.eval) {
			return false
		}
		if p.eval[p.idx] != ' ' {
			break
		}
		p.idx++
	}

	if p.eval[p.idx] >= '0' && p.eval[p.idx] <= '9' {
		var value int
		for p.idx < len(p.eval) {
			if p.eval[p.idx] >= '0' && p.eval[p.idx] <= '9' {
				value *= 10
				value += int(p.eval[p.idx] - '0')
				p.idx++
			} else {
				break
			}
		}
		p.parsedToken = number(value)
		return true
	} else if p.eval[p.idx] == '+' {
		p.parsedToken = plus{}
		p.idx++
		return true
	} else if p.eval[p.idx] == '-' {
		p.parsedToken = minus{}
		p.idx++
		return true
	} else if p.eval[p.idx] == '*' {
		p.parsedToken = multiple{}
		p.idx++
		return true
	} else if p.eval[p.idx] == '/' {
		p.parsedToken = divide{}
		p.idx++
		return true
	}

	return false
}

func tokenize(eval string) []token {
	tokens := []token{}
	p := &parser{
		eval: []rune(eval),
	}
	for p.parse() {
		tokens = append(tokens, p.parsedToken)
	}
	return tokens
}

func postfix(eval string) []token {
	tokens := tokenize(eval)
	if len(tokens) == 0 {
		return tokens
	}

	postfix := make([]token, 0, len(tokens))
	ops := []token{}

	for i := 0; i < len(tokens); i++ {
		if no, ok := tokens[i].(number); ok {
			postfix = append(postfix, no)
		} else {
			for len(ops) > 0 && ops[len(ops)-1].Priority() >= tokens[i].Priority() {
				postfix = append(postfix, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, tokens[i])
		}
	}
	for len(ops) > 0 {
		postfix = append(postfix, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}
	return postfix
}

func Evaluate(eval string) (rst int, success bool) {
	tokens := postfix(eval)

	top, tokens := tokens[len(tokens)-1], tokens[:len(tokens)-1]
	rst, success = top.Evaluate(&tokens)
	return
}

func MakeExpressionTree(eval string) (*binarytree.TreeNode, bool) {
	tokens := postfix(eval)

	top, tokens := tokens[len(tokens)-1], tokens[:len(tokens)-1]
	return top.MakeTreeNode(&tokens)
}

func EvaluateExpressionTree(root *binarytree.TreeNode) int {
	t := root.Value.(token)
	if t == nil {
		return 0
	}

	return t.EvaluateTree(root.Left, root.Right)
}
