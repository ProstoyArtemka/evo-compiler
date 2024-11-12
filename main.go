package main

import (
	"evo-compiler/src/interpreter"
	"evo-compiler/src/lexer"
	"evo-compiler/src/parser"
	"os"
)

func readFile(filePath string) string {

	file, err := os.ReadFile(filePath)

	if err != nil {
		panic("Error while reading a file")
	}

	return string(file)
}

func main() {
	if len(os.Args) < 1 {
		return
	}

	var filePath string = os.Args[1]
	var fileContent = readFile(filePath)

	var tokens []lexer.Token = lexer.ToTokens(fileContent)

	// for i := 0; i < len(tokens); i++ {
	// 	fmt.Printf("Position: %v\n", i)

	// 	tokens[i].Print()

	// 	fmt.Println()
	// }

	var ast []parser.Node = parser.ParseTokens(tokens)
	interpreter.Run(ast)
}
