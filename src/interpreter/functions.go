package interpreter

import (
	"evo-compiler/src/lexer"
	"evo-compiler/src/parser"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Functions map[string]func([]any, Scope) any

type DeclaredFunction struct {
	Arguments   []lexer.Token
	Expressions []parser.Node
}

var GlobalFunctions Functions = Functions{
	"print":   Print,
	"println": PrintLn,

	"read":   Read,
	"random": Random,

	"num": Num,
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

func PrintLn(args []any, scope Scope) any {
	return Print(append(args, "\n"), scope)
}

func Read(args []any, scope Scope) any {
	var ret = ""

	fmt.Scanf("%s", &ret)

	return ret
}

func Random(args []any, scope Scope) any {
	rand.Seed(time.Now().UnixNano())

	if len(args) == 1 {
		var to = GetUnConstantValue(args[0])
		var toType = GetType(to)

		if toType == INTEGER {
			return rand.Intn(to.(int))
		}

		if toType == FLOAT {
			return rand.Float64() * to.(float64)
		}
	}

	if len(args) == 2 {
		var from = GetUnConstantValue(args[0])
		var to = GetUnConstantValue(args[1])

		var fromType = GetType(from)
		var toType = GetType(to)

		if fromType == INTEGER {
			from = float64(from.(int))
		}

		if toType == INTEGER {
			to = float64(to.(int))
		}

		var result = (rand.Float64() * (to.(float64) - from.(float64))) + from.(float64)

		if toType == INTEGER && fromType == INTEGER {
			return int(result)
		}

		return result
	}

	return 0
}

func Num(args []any, scope Scope) any {

	if len(args) == 0 {
		return 0
	}

	var arg = args[0]

	if str, ok := arg.(string); ok {
		res, err := strconv.Atoi(str)

		if err != nil {
			return -1
		}

		return res
	}

	return 0
}
