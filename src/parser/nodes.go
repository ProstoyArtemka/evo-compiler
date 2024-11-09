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

type VariableNode struct {
	Node

	Name lexer.Token
}

type ConstantNode struct {
	Node

	Value lexer.Token
}

var nilToken = lexer.Token{}
