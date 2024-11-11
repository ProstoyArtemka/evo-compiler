package interpreter

import (
	"evo-compiler/src/lexer"
	"evo-compiler/src/parser"
	"fmt"
)

type Functions map[string]func([]any, Scope) any

type DeclaredFunction struct {
	Arguments   []lexer.Token
	Expressions []parser.Node
}

var GlobalFunctions Functions = Functions{
	"print": Print,
}

func Print(args []any, scope Scope) any {
	var ret = ""

	for i := 0; i < len(args); i++ {
		var arg = GetUnConstantValue(args[i])
		var space = ""

		if i < len(args)-1 {
			space = " "
		}

		fmt.Print(arg, space)
		ret += fmt.Sprint(arg) + space
	}

	return ret
}
