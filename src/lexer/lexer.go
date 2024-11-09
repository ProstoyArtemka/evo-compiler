package lexer

import (
	"regexp"
)

var lexem string = ""
var tokens []Token = make([]Token, 0)

var spacesRegex *regexp.Regexp = regexp.MustCompile(`^[ \n\t\r]$`)
var operatorsRegex *regexp.Regexp = regexp.MustCompile(`^[\=\+\-\*\/\(\)\!]$`)

var namesRegex *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z_]+[a-zA-Z_0-9]*$`)

var integerRegex *regexp.Regexp = regexp.MustCompile(`^[0-9]$`)
var floatRegex *regexp.Regexp = regexp.MustCompile(`^[0-9]+.[0-9]+$`)

func PushLexem() {
	if lexem == "" {
		return
	}

	var tokenType int = -1

	if lexem == ";" {
		tokenType = NEWLINE
	}

	// OPERATORS

	if BINARY_OPERATORS.Contains(lexem) {
		tokenType = BINARY_OPERATOR
	}

	if lexem == "=" {
		tokenType = ASSIGN
	}

	// NAMES

	if namesRegex.MatchString(lexem) {
		tokenType = NAME
	}

	// CONSTANTS

	if integerRegex.MatchString(lexem) {
		tokenType = INTEGER
	}

	if floatRegex.MatchString(lexem) {
		tokenType = FLOAT
	}

	// BRACKETS

	if lexem == "(" {
		tokenType = L_BRACKET
	}

	if lexem == ")" {
		tokenType = R_BRACKET
	}

	var token = Token{Lexem: lexem, TokenType: tokenType}
	tokens = append(tokens, token)

	lexem = ""
}

func ToTokens(content string) []Token {
	for pos := 0; pos < len(content); pos++ {
		var current = string(content[pos])

		if operatorsRegex.MatchString(current) {
			PushLexem()

			lexem += current

			var next = string(content[pos+1])
			for len(content) < (pos+1) && operatorsRegex.MatchString(next) && OPERATORS.Contains(current+next) && len(lexem) < 2 {
				lexem += string(content[pos+1])

				next = string(content[pos+1])
			}

			PushLexem()

			continue
		}

		if current == ";" {
			PushLexem()

			lexem = ";"

			PushLexem()

			continue
		}

		if spacesRegex.MatchString(current) {
			PushLexem()

			continue
		}

		lexem += current
	}

	PushLexem()

	return tokens
}
