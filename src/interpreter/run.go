package interpreter

import (
	"evo-compiler/src/parser"
	"fmt"
)

type ConstantValue struct {
	value        string
	constantType int
}

var scope map[string]any = make(map[string]any)

func RunBinaryOperator(node parser.BinaryOperatorNode) any {

	var leftValue = GetValueFromNode(node.Left)
	var rightValue = GetValueFromNode(node.Right)

	switch node.Operator.Lexem {

	case "+":
		return SumOf(leftValue, rightValue)

	case "-":
		return SubOf(leftValue, rightValue)

	}

	return nil
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

	return nil
}

func RunAssignNode(node parser.AssignNode) {

	scope[node.Left.Name.Lexem] = GetValueFromNode(node.Right)

}

func RunNode(node parser.Node) {

	if assignNode, ok := node.(parser.AssignNode); ok {
		RunAssignNode(assignNode)
	}

}

func Run(nodes []parser.Node) {

	for i := 0; i < len(nodes); i++ {
		RunNode(nodes[i])
	}

	fmt.Println(scope)
}
