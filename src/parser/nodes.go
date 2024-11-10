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

type GlobalAssignNode struct {
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
	Else       Node
}

type ElseNode struct {
	Node

	Expressions []Node
}

type CallFunctionNode struct {
	Node

	Name      lexer.Token
	Arguments []Node
}

var NullToken = lexer.Token{}
var NullNode = NilNode{}
