package nodeinterface

type Node interface {
	GetChilds() []Node
	GetValue() any
}
