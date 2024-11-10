package parser

import (
	"evo-compiler/src/lexer"
	"evo-compiler/src/utils"
	"fmt"
	"reflect"
)

var tokens []lexer.Token = nil
var expressions []Node = make([]Node, 0)
var position = 0

var EXPRESSIONS_WITHOUT_NEWLINE []Node = []Node{IfNode{}}

func IsWithoutNewline(expression Node) bool {

	for i := 0; i < len(EXPRESSIONS_WITHOUT_NEWLINE); i++ {
		var expr = EXPRESSIONS_WITHOUT_NEWLINE[i]

		if reflect.TypeOf(expr) == reflect.TypeOf(expression) {
			return true
		}

	}

	return false
}

func Match(types ...int) lexer.Token {

	if position < len(tokens) {
		var currentToken lexer.Token = tokens[position]

		if utils.IntContains(types, currentToken.TokenType) {
			position += 1

			return currentToken
		}
	}

	return nilToken
}

func Expect(types ...int) {

	if Match(types...) == nilToken {
		panic(fmt.Sprintf("Expected other value in position %v", position))
	}
}

func ParseConstantOrVariable() Node {
	var constant = Match(lexer.INTEGER, lexer.FLOAT, lexer.STRING, lexer.FALSE, lexer.TRUE)

	if constant != nilToken {
		return ConstantNode{Value: constant}
	}

	var name = Match(lexer.NAME)

	if name != nilToken {
		return VariableNode{Name: name}
	}

	return NilNode{}
}

func ParseBrackets() Node {

	if Match(lexer.L_BRACKET) != nilToken {
		var node = ParseFormula()

		Expect(lexer.R_BRACKET)

		return node
	}

	return ParseConstantOrVariable()
}

func ParseFormula() Node {
	var left = ParseBrackets()

	if left == nilNode {
		var operator = Match(lexer.UNARY_OPERATOR)

		if operator != nilToken {
			var right = ParseBrackets()

			return UnaryOperatorNode{Operand: right, Operator: operator}
		}
	}

	var operator = Match(lexer.BINARY_OPERATOR)

	for operator != nilToken {
		var right = ParseBrackets()
		left = BinaryOperatorNode{Left: left, Right: right, Operator: operator}

		operator = Match(lexer.BINARY_OPERATOR)
	}

	return left
}

func ParseExpression() Node {

	if Match(lexer.NAME) != nilToken {

		if Match(lexer.ASSIGN) != nilToken {
			position -= 2

			var name lexer.Token = Match(lexer.NAME)
			var variable = VariableNode{Name: name}

			Expect(lexer.ASSIGN)

			var node = ParseFormula()

			return AssignNode{Left: variable, Right: node}
		}
	}

	if Match(lexer.IF) != nilToken {
		var formula = ParseFormula()

		Expect(lexer.C_L_BRACKET)

		var expressions []Node = make([]Node, 0)
		for Match(lexer.C_R_BRACKET) == nilToken {
			var expression = ParseExpression()
			Expect(lexer.NEWLINE)

			expressions = append(expressions, expression)
		}

		return IfNode{Formula: formula, Expresions: expressions}
	}

	return NilNode{}
}

func ParseTokens(inTokens []lexer.Token) []Node {
	tokens = inTokens

	for position < len(tokens) {
		var expression = ParseExpression()

		expressions = append(expressions, expression)

		var isWithoutNewline = IsWithoutNewline(expression)

		if !isWithoutNewline {
			Expect(lexer.NEWLINE)
		}
	}

	return expressions
}
