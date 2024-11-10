package parser

import "evo-compiler/src/lexer"

type Node interface{}

type NilNode struct {
	Node
}

type AssignNode struct {
	Node

	Left  VariableNode
	Right Node
}

type BinaryOperatorNode struct {
	Node

	Left     Node
	Right    Node
	Operator lexer.Token
}

type UnaryOperatorNode struct {
	Node

	Operand  Node
	Operator lexer.Token
}

type VariableNode struct {
	Node

	Name lexer.Token
}

type ConstantNode struct {
	Node

	Value lexer.Token
}

type IfNode struct {
	Node

	Formula    Node
	Expresions []Node
}

var nilToken = lexer.Token{}
var nilNode = NilNode{}
