package main

import (
	"evo-compiler/src/interpreter"
	"evo-compiler/src/lexer"
	"evo-compiler/src/parser"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func readFile(filePath string) string {
	codec := unicode.UTF8

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	rd := transform.NewReader(file, codec.NewDecoder())
	data, err := ioutil.ReadAll(rd)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	if len(os.Args) < 1 {
		return
	}

	var filePath string = os.Args[1]
	var fileContent = readFile(filePath)

	var tokens []lexer.Token = lexer.ToTokens(fileContent)

	for i := 0; i < len(tokens); i++ {
		fmt.Printf("Position: %v\n", i)

		tokens[i].Print()

		fmt.Println()
	}

	var ast []parser.Node = parser.ParseTokens(tokens)
	interpreter.Run(ast)
}
