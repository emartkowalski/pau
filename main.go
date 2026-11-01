package main

import (
	"path/filepath"
	"strings"
	"fmt"
	"os"
	"pau/src/syntax"
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
		generate_HTML_File(subcomandArgs)
	case "-h", "-help", "help":
		fmt.Println("=== help ===")
		fmt.Println("run 'example.p'")
		os.Exit(0)
	default:
		fmt.Printf("uknown command %q\n", subcomand)
		os.Exit(1)
	}

}

func generate_HTML_File(filepathPau string) {

	content, err := os.ReadFile(filepathPau)
	if err != nil {
		fmt.Printf("read file failed: %v\n", err)
		os.Exit(1)
	}
	lexer := syntax.NewLexer(string(content))
	parser := syntax.NewParser(lexer, filepathPau)
	nodes, err := parser.Parse()
	if err != nil {
		fmt.Println("#pau")
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// for _, node := range nodes {
	// 	fmt.Print(node.String())
	// }
	ext := filepath.Ext(filepathPau)
	newfilepathPau := strings.TrimSuffix(filepathPau, ext) + ".html"

	file, err := os.Create(newfilepathPau)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, node := range nodes {
		_, err := file.WriteString(node.String())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Printf("done! %q\n", newfilepathPau)

}

/*
`

*/
