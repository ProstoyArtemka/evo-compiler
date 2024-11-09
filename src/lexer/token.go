package lexer

import (
	"evo-compiler/src/utils"
	"fmt"
)

const (
	ASSIGN          = 0
	BINARY_OPERATOR = 1

	NAME = 2

	INTEGER = 3
	FLOAT   = 4
	STRING  = 5
	NEWLINE = 6

	L_BRACKET = 7
	R_BRACKET = 8
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
