package parser

import (
	"evo-compiler/src/lexer"
	"evo-compiler/src/utils"
	"fmt"
)

var tokens []lexer.Token = nil
var expressions []Node = make([]Node, 0)
var position = 0

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
	var constant = Match(lexer.INTEGER, lexer.FLOAT, lexer.STRING)

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

	return NilNode{}
}

func ParseTokens(inTokens []lexer.Token) []Node {
	tokens = inTokens

	for position < len(tokens) {
		var expression = ParseExpression()

		expressions = append(expressions, expression)

		Expect(lexer.NEWLINE)
	}

	return expressions
}
