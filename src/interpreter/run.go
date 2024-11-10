package interpreter

import (
	"evo-compiler/src/parser"
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

type Scope map[string]any

var GlobalScope Scope = make(Scope, 0)

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

func RunBinaryOperator(node parser.BinaryOperatorNode, scope Scope) any {

	var leftValue = GetValueFromNode(node.Left, scope)
	var rightValue = GetValueFromNode(node.Right, scope)

	var binaryFn = BinaryOperatorFunctions[node.Operator.Lexem]

	if binaryFn == nil {
		return nil
	}

	return binaryFn(leftValue, rightValue)
}

func RunUnaryOperator(node parser.UnaryOperatorNode, scope Scope) any {

	var operand = GetValueFromNode(node.Operand, scope)
	var unaryFn = UnaryOperatorFunctions[node.Operator.Lexem]

	if unaryFn == nil {
		return nil
	}

	return unaryFn(operand)
}

func GetValueFromNode(node parser.Node, scope Scope) any {

	if constantNode, ok := node.(parser.ConstantNode); ok {
		return ConstantValue{value: constantNode.Value.Lexem, constantType: constantNode.Value.TokenType}
	}

	if variableNode, ok := node.(parser.VariableNode); ok {
		return scope[variableNode.Name.Lexem]
	}

	if binaryNode, ok := node.(parser.BinaryOperatorNode); ok {
		return RunBinaryOperator(binaryNode, scope)
	}

	if unaryNode, ok := node.(parser.UnaryOperatorNode); ok {
		return RunUnaryOperator(unaryNode, scope)
	}

	return nil
}

func RunAssignNode(node parser.AssignNode, scope Scope) {

	scope[node.Left.Name.Lexem] = GetValueFromNode(node.Right, scope)

}

func RunGlobalAssignNode(node parser.GlobalAssignNode) {

	GlobalScope[node.Left.Name.Lexem] = GetValueFromNode(node.Right, GlobalScope)

}

func RunIfNode(node parser.IfNode, scope Scope) {

	var formula = GetValueFromNode(node.Formula, scope)

	if formula == true {
		var localScope = make(Scope, 0)

		for i := 0; i < len(node.Expresions); i++ {
			var node = node.Expresions[i]

			RunNode(node, localScope)
		}
	} else if formula == false && node.Else != parser.NullNode {
		var elseNode = node.Else.(parser.ElseNode)
		var localScope = make(Scope, 0)

		for i := 0; i < len(elseNode.Expressions); i++ {
			var expression = elseNode.Expressions[i]

			RunNode(expression, localScope)
		}
	}
}

func RunCallFunction(node parser.CallFunctionNode, scope Scope) {
	var callArgs []any = make([]any, 0)

	for i := 0; i < len(node.Arguments); i++ {
		var val = GetValueFromNode(node.Arguments[i], scope)

		callArgs = append(callArgs, val)
	}

	var fn = GlobalFunctions[node.Name.Lexem]

	if fn != nil {
		fn(callArgs, scope)
	}
}

func RunNode(node parser.Node, scope Scope) {
	var currentScope = scope

	if scope == nil {
		currentScope = GlobalScope
	}

	if assignNode, ok := node.(parser.AssignNode); ok {
		RunAssignNode(assignNode, currentScope)
	}

	if globalAssignNode, ok := node.(parser.GlobalAssignNode); ok {
		RunGlobalAssignNode(globalAssignNode)
	}

	if ifNode, ok := node.(parser.IfNode); ok {
		RunIfNode(ifNode, currentScope)
	}

	if callFuncNode, ok := node.(parser.CallFunctionNode); ok {
		RunCallFunction(callFuncNode, scope)
	}

}

func Run(nodes []parser.Node) {

	for i := 0; i < len(nodes); i++ {
		RunNode(nodes[i], nil)
	}
}
