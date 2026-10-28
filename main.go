package main

import (
	"fmt"
	"pau/src/parser_1"
	"os"
)

func main() {
	filePath := "main.lt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file failed: %v\n", err)
		os.Exit(1)
	}
	lexer := parser_1.NewLexer(string(content))
	parser := parser_1.NewParser(lexer, filePath)
	nodes, err := parser.Parse()
	if err != nil {
		fmt.Println("#pau")
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	for _, node := range nodes {
		fmt.Print(node.String())
	}
}
