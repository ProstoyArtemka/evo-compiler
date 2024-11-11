package lexer

import (
	"regexp"
	"strings"
)

var lexem string = ""
var tokens []Token = make([]Token, 0)

var spacesRegex *regexp.Regexp = regexp.MustCompile(`^[ \n\t\r]$`)
var operatorsRegex *regexp.Regexp = regexp.MustCompile(`^[\=\+\-\*\/\(\)\!\&\|\{\}\:\,\>\<]$`)

var namesRegex *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z_]+[a-zA-Z_0-9]*$`)

var integerRegex *regexp.Regexp = regexp.MustCompile(`^[0-9]$`)
var floatRegex *regexp.Regexp = regexp.MustCompile(`^[0-9]+.[0-9]+$`)

func PushLexem(tokenType int) {
	if lexem == "" {
		return
	}

	if tokenType == -1 {

		if namesRegex.MatchString(lexem) {
			tokenType = NAME
		}

		if KEYWORDS.Contains(lexem) {
			tokenType = KEYWORD_TYPES[lexem]
		}

		if integerRegex.MatchString(lexem) {
			tokenType = INTEGER
		}

		if floatRegex.MatchString(lexem) {
			tokenType = FLOAT
		}

	}

	var token = Token{Lexem: lexem, TokenType: tokenType}
	tokens = append(tokens, token)

	lexem = ""
}

func ToTokens(content string) []Token {
	for pos := 0; pos < len(content); pos++ {
		var current = string(content[pos])

		if current == "\"" && len(content) > (pos+1) {
			var next = string(content[pos+1])

			for pos < len(content) && next != "\"" {
				lexem += next

				pos++
				next = string(content[pos+1])
			}

			pos++

			lexem = strings.ReplaceAll(lexem, "\\n", "\n")
			lexem = strings.ReplaceAll(lexem, "\\t", "\t")

			PushLexem(STRING)

			continue
		}

		if operatorsRegex.MatchString(current) {
			PushLexem(-1)

			lexem += current

			if pos+1 < len(content) {
				var next = string(content[pos+1])

				for pos < len(content) && operatorsRegex.MatchString(next) && OPERATORS.Contains(current+next) && len(lexem) < 2 {
					lexem += string(content[pos+1])

					pos++

					next = string(content[pos+1])
				}
			}

			if BINARY_OPERATORS.Contains(lexem) {
				PushLexem(BINARY_OPERATOR)

			} else if UNARY_OPERATORS.Contains(lexem) {
				PushLexem(UNARY_OPERATOR)

			} else {
				PushLexem(OPERATORS_TYPES[lexem])

			}

			continue
		}

		if current == ";" {
			PushLexem(-1)

			lexem = ";"

			PushLexem(NEWLINE)

			continue
		}

		if spacesRegex.MatchString(current) {
			PushLexem(-1)

			continue
		}

		lexem += current
	}

	PushLexem(-1)

	return tokens
}
