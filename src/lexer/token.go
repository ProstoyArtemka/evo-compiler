package lexer

import (
	"evo-compiler/src/utils"
	"fmt"
)

const (
	INVALID = -1

	ASSIGN          = 0
	BINARY_OPERATOR = 1
	OPERATOR        = 2

	NAME = 3

	INTEGER = 4
	FLOAT   = 5
	STRING  = 6
	NEWLINE = 7

	L_BRACKET = 8
	R_BRACKET = 9

	EOF = 256
)

type Token struct {
	Lexem     string
	TokenType int
}

func (token Token) Print() {
	fmt.Printf("Lexem: %s\nType: %s\n", token.Lexem, fmt.Sprint(token.TokenType))
}

var OPERATORS = utils.StringArray{"-", "+", "=", "*", "/", "!", "==", "!=", "(", ")"}
var BINARY_OPERATORS = utils.StringArray{"-", "+", "*", "/", "==", "!="}

var OPERATORS_TYPES map[string]int = map[string]int{
	"(": L_BRACKET,
	")": R_BRACKET,
	"=": ASSIGN,
}
