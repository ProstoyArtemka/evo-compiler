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

var EXPRESSIONS_WITHOUT_NEWLINE []Node = []Node{IfNode{}, DeclareFunctionNode{}}

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

	return NullToken
}

func Expect(types ...int) {

	if Match(types...) == NullToken {
		panic(fmt.Sprintf("Expected other value in position %v", position))
	}
}

func ParseConstantOrVariable() Node {
	var constant = Match(lexer.INTEGER, lexer.FLOAT, lexer.STRING, lexer.FALSE, lexer.TRUE)

	if constant != NullToken {
		return ConstantNode{Value: constant}
	}

	var name = Match(lexer.NAME)

	if Match(lexer.L_BRACKET) != NullToken {

		var arguments []Node
		for Match(lexer.R_BRACKET) == NullToken {
			var arg = ParseFormula()

			arguments = append(arguments, arg)

			if Match(lexer.COMMA) == NullToken {
				break
			}
		}

		position++

		return CallFunctionNode{Name: name, Arguments: arguments}
	}

	if name != NullToken {
		return VariableNode{Name: name}
	}

	return NilNode{}
}

func ParseBrackets() Node {

	if Match(lexer.L_BRACKET) != NullToken {
		var node = ParseFormula()

		Expect(lexer.R_BRACKET)

		return node
	}

	return ParseConstantOrVariable()
}

func ParseFormula() Node {
	var left = ParseBrackets()

	if left == NullNode {
		var operator = Match(lexer.UNARY_OPERATOR)

		if operator != NullToken {
			var right = ParseBrackets()

			return UnaryOperatorNode{Operand: right, Operator: operator}
		}
	}

	var operator = Match(lexer.BINARY_OPERATOR)

	for operator != NullToken {
		var right = ParseBrackets()
		left = BinaryOperatorNode{Left: left, Right: right, Operator: operator}

		operator = Match(lexer.BINARY_OPERATOR)
	}

	if operator == NullToken && Match(lexer.QUESTION_MARK) != NullToken {
		var trueFormula = ParseFormula()

		Expect(lexer.COLON)

		var falseFormula = ParseFormula()

		return TernaryOperator{BoolExpression: left, TrueExpression: trueFormula, FalseExpression: falseFormula}
	}

	return left
}

func ParseExpressionsUntilBracket() []Node {
	var expressions = make([]Node, 0)

	for Match(lexer.C_R_BRACKET) == NullToken {
		var expression = ParseExpression()

		if !IsWithoutNewline(expression) {
			Expect(lexer.NEWLINE)
		}

		expressions = append(expressions, expression)
	}

	return expressions
}

func ParseExpression() Node {

	if name := Match(lexer.NAME); name != NullToken {

		if assign := Match(lexer.ASSIGN, lexer.GLOBAL_ASSIGN); assign != NullToken {
			position -= 2

			var name lexer.Token = Match(lexer.NAME)
			var variable = VariableNode{Name: name}

			Expect(assign.TokenType)

			var node = ParseFormula()

			if assign.TokenType == lexer.ASSIGN {
				return AssignNode{Left: variable, Right: node}
			} else {
				return GlobalAssignNode{Left: variable, Right: node}
			}
		}

		if Match(lexer.L_BRACKET) != NullToken {

			var arguments []Node
			for Match(lexer.R_BRACKET) == NullToken {
				var arg = ParseFormula()

				arguments = append(arguments, arg)

				if Match(lexer.COMMA) == NullToken {
					Expect(lexer.R_BRACKET)

					break
				}
			}

			return CallFunctionNode{Name: name, Arguments: arguments}
		}
	}

	if Match(lexer.IF) != NullToken {
		var formula = ParseFormula()

		Expect(lexer.C_L_BRACKET)

		var expressions []Node = ParseExpressionsUntilBracket()

		var elseExpression Node = NullNode

		if Match(lexer.ELSE) != NullToken {
			Expect(lexer.C_L_BRACKET)
			var elseExpressions []Node = ParseExpressionsUntilBracket()

			elseExpression = ElseNode{Expressions: elseExpressions}
		}

		return IfNode{Formula: formula, Expresions: expressions, Else: elseExpression}
	}

	if Match(lexer.FUNCTION) != NullToken {
		var name = Match(lexer.NAME)

		Expect(lexer.L_BRACKET)

		var args []lexer.Token = make([]lexer.Token, 0)
		for Match(lexer.R_BRACKET) == NullToken {
			var arg = Match(lexer.NAME)

			if arg == NullToken {
				Expect(lexer.R_BRACKET)

				break
			}

			args = append(args, arg)

			if Match(lexer.COMMA) == NullToken {
				Expect(lexer.R_BRACKET)

				break
			}
		}

		Expect(lexer.C_L_BRACKET)

		var expressions []Node = ParseExpressionsUntilBracket()

		return DeclareFunctionNode{Name: name, Expressions: expressions, Arguments: args}
	}

	if (Match(lexer.RETURN)) != NullToken {
		var formula = ParseFormula()

		return ReturnNode{Value: formula}
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
