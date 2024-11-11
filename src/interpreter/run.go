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

const (
	OK       = 0
	RETURNED = 1
)

type State int

type Scope map[string]any

var GlobalScope Scope = make(Scope, 0)

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
		var scopeVar = scope[variableNode.Name.Lexem]
		var globalVar = GlobalScope[variableNode.Name.Lexem]

		if scopeVar != nil {
			return scopeVar
		}

		if globalVar != nil {
			return globalVar
		}

		return nil
	}

	if binaryNode, ok := node.(parser.BinaryOperatorNode); ok {
		return RunBinaryOperator(binaryNode, scope)
	}

	if unaryNode, ok := node.(parser.UnaryOperatorNode); ok {
		return RunUnaryOperator(unaryNode, scope)
	}

	if callNode, ok := node.(parser.CallFunctionNode); ok {
		return RunCallFunction(callNode, scope)
	}

	if returnNode, ok := node.(parser.ReturnNode); ok {
		return GetValueFromNode(returnNode.Value, scope)
	}

	return nil
}

func RunAssignNode(node parser.AssignNode, scope Scope) {

	scope[node.Left.Name.Lexem] = GetValueFromNode(node.Right, scope)

}

func RunGlobalAssignNode(node parser.GlobalAssignNode) {

	GlobalScope[node.Left.Name.Lexem] = GetValueFromNode(node.Right, GlobalScope)

}

func RunIfNode(node parser.IfNode, scope Scope) (State, any) {

	var formula = GetValueFromNode(node.Formula, scope)

	if formula == true {
		var localScope = make(Scope, 0)

		for i := 0; i < len(node.Expresions); i++ {
			var node = node.Expresions[i]

			state, val := RunNode(node, localScope)

			if state != OK {
				return state, val
			}
		}

	} else if formula == false && node.Else != parser.NullNode {
		var elseNode = node.Else.(parser.ElseNode)
		var localScope = make(Scope, 0)

		for i := 0; i < len(elseNode.Expressions); i++ {
			var expression = elseNode.Expressions[i]

			state, val := RunNode(expression, localScope)

			if state != OK {
				return state, val
			}
		}
	}

	return OK, nil
}

func RunDeclaredFunction(function DeclaredFunction, args []any, scope Scope) any {
	var returnValue any = nil
	var localScope = make(Scope, 0)

	for i := 0; i < len(args); i++ {
		var arg = args[i]
		var argName = function.Arguments[i]

		localScope[argName.Lexem] = arg
	}

	for i := 0; i < len(function.Expressions); i++ {
		var expression = function.Expressions[i]

		if returnNode, ok := expression.(parser.ReturnNode); ok {
			return GetValueFromNode(returnNode, localScope)
		}

		state, val := RunNode(expression, localScope)

		if state == RETURNED {
			return val
		}
	}

	return returnValue
}

func RunCallFunction(node parser.CallFunctionNode, scope Scope) any {
	var callArgs []any = make([]any, 0)

	for i := 0; i < len(node.Arguments); i++ {
		var val = GetValueFromNode(node.Arguments[i], scope)

		callArgs = append(callArgs, val)
	}

	var globalFn = GlobalFunctions[node.Name.Lexem]
	var scopeFn = scope[node.Name.Lexem]
	var globalScopeFn = GlobalScope[node.Name.Lexem]

	if globalFn != nil {
		return globalFn(callArgs, scope)
	}

	declared, ok := scopeFn.(DeclaredFunction)
	declaredGlobal, globalOk := globalScopeFn.(DeclaredFunction)

	if scopeFn != nil && ok {
		return RunDeclaredFunction(declared, callArgs, scope)
	}

	if globalScopeFn != nil && globalOk {
		return RunDeclaredFunction(declaredGlobal, callArgs, scope)
	}

	return nil
}

func RunDeclareFunction(node parser.DeclareFunctionNode, scope Scope) any {
	var declared = DeclaredFunction{Expressions: node.Expressions, Arguments: node.Arguments}

	scope[node.Name.Lexem] = declared

	return declared
}

func RunNode(node parser.Node, scope Scope) (State, any) {
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
		return RunIfNode(ifNode, currentScope)
	}

	if callFuncNode, ok := node.(parser.CallFunctionNode); ok {
		RunCallFunction(callFuncNode, currentScope)
	}

	if declareFuncNode, ok := node.(parser.DeclareFunctionNode); ok {
		RunDeclareFunction(declareFuncNode, currentScope)
	}

	if returnNode, ok := node.(parser.ReturnNode); ok {
		return RETURNED, GetValueFromNode(returnNode, currentScope)
	}

	return OK, nil
}

func Run(nodes []parser.Node) {

	for i := 0; i < len(nodes); i++ {
		state, _ := RunNode(nodes[i], nil)

		if state == RETURNED {
			break
		}
	}
}
