package interpreter

import (
	"evo-compiler/src/parser"
	"fmt"
)

type ConstantValue struct {
	value        string
	constantType int
}

const (
	INTEGER = 0
	FLOAT   = 1
	STRING  = 2
	BOOLEAN = 3
)

var scope map[string]any = make(map[string]any)

var BinaryOperatorFunctions map[string]func(left any, right any) any = map[string]func(left any, right any) any{
	"+":  SumOf,
	"-":  SubOf,
	"*":  MulOf,
	"/":  DivOf,
	"&&": AndOf,
	"||": OrOf,
	"==": Equals,
	"!=": NotEquals,
}

var UnaryOperatorFunctions map[string]func(operand any) any = map[string]func(operand any) any{
	"!": Not,
}

func RunBinaryOperator(node parser.BinaryOperatorNode) any {

	var leftValue = GetValueFromNode(node.Left)
	var rightValue = GetValueFromNode(node.Right)

	var binaryFn = BinaryOperatorFunctions[node.Operator.Lexem]

	if binaryFn == nil {
		return nil
	}

	return binaryFn(leftValue, rightValue)
}

func RunUnaryOperator(node parser.UnaryOperatorNode) any {

	var operand = GetValueFromNode(node.Operand)
	var unaryFn = UnaryOperatorFunctions[node.Operator.Lexem]

	if unaryFn == nil {
		return nil
	}

	return unaryFn(operand)
}

func GetValueFromNode(node parser.Node) any {

	if constantNode, ok := node.(parser.ConstantNode); ok {
		return ConstantValue{value: constantNode.Value.Lexem, constantType: constantNode.Value.TokenType}
	}

	if variableNode, ok := node.(parser.VariableNode); ok {
		return scope[variableNode.Name.Lexem]
	}

	if binaryNode, ok := node.(parser.BinaryOperatorNode); ok {
		return RunBinaryOperator(binaryNode)
	}

	if unaryNode, ok := node.(parser.UnaryOperatorNode); ok {
		return RunUnaryOperator(unaryNode)
	}

	return nil
}

func RunAssignNode(node parser.AssignNode) {

	scope[node.Left.Name.Lexem] = GetValueFromNode(node.Right)

}

func RunIfNode(node parser.IfNode) {

	var formula = GetValueFromNode(node.Formula)

	if formula == true {

		for i := 0; i < len(node.Expresions); i++ {
			var node = node.Expresions[i]

			RunNode(node)
		}
	}
}

func RunNode(node parser.Node) {

	if assignNode, ok := node.(parser.AssignNode); ok {
		RunAssignNode(assignNode)
	}

	if ifNode, ok := node.(parser.IfNode); ok {
		RunIfNode(ifNode)
	}

}

func Run(nodes []parser.Node) {

	for i := 0; i < len(nodes); i++ {
		RunNode(nodes[i])
	}

	fmt.Println(scope)
}
