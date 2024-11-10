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

	C_L_BRACKET = 10
	C_R_BRACKET = 11

	IF    = 12
	TRUE  = 13
	FALSE = 14

	UNARY_OPERATOR = 15

	EOF = 256
)

type Token struct {
	Lexem     string
	TokenType int
}

func (token Token) Print() {
	fmt.Printf("Lexem: %s\nType: %s\n", token.Lexem, fmt.Sprint(token.TokenType))
}

var OPERATORS = utils.StringArray{"-", "+", "=", "*", "/", "!", "==", "!=", "(", ")", "{", "}", "&", "|", "&&", "||"}
var BINARY_OPERATORS = utils.StringArray{"-", "+", "*", "/", "==", "!=", "&&", "||"}

var UNARY_OPERATORS = utils.StringArray{"!"}

var OPERATORS_TYPES map[string]int = map[string]int{
	"(": L_BRACKET,
	")": R_BRACKET,

	"{": C_L_BRACKET,
	"}": C_R_BRACKET,

	"=": ASSIGN,
}

var KEYWORDS = utils.StringArray{"if", "true", "false"}
var KEYWORD_TYPES map[string]int = map[string]int{
	"if":    IF,
	"true":  TRUE,
	"false": FALSE,
}
