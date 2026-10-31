package main

import (
	"fmt"
	"pau/src/parser_1"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("not enouth arguments\n")
		os.Exit(1)
	}

	// subcomand := os.Args[0]
	subcomand := os.Args[1]

	switch subcomand {
	case "run":
		if len(os.Args) < 3 {
			fmt.Printf("not enouth arguments\n")
			os.Exit(1)
		}

		subcomandArgs := os.Args[2]
		doWork(subcomandArgs)
	case "-h", "-help", "help":
		fmt.Println("=== help ===")
		fmt.Println("run 'example.p'")
		os.Exit(0)
	default:
		fmt.Printf("uknown command %q\n", subcomand)
		os.Exit(1)
	}

}

func doWork(filepath string) {

	filePath := filepath
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

/*
`

*/
