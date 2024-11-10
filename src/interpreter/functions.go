package interpreter

import "fmt"

type Functions map[string]func([]any, Scope) any

var GlobalFunctions Functions = Functions{
	"print": Print,
}

func Print(args []any, scope Scope) any {
	for i := 0; i < len(args); i++ {
		var arg = GetUnConstantValue(args[i])
		var space = ""

		if i < len(args)-1 {
			space = " "
		}

		fmt.Print(arg, space)
	}

	return nil
}
