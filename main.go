package main

import (
	"path/filepath"
	"strings"
	"io/fs"
	"fmt"
	"os"
	"pau/src/syntax"
)

var pauVersion string = "v0.2.0"

func main() {

	if len(os.Args) < 2 {
		logo()
		fmt.Printf("not enouth arguments\n")
		os.Exit(1)
	}

	// subcomand := os.Args[0]
	subcomand := os.Args[1]

	switch subcomand {
	case "run":
		if len(os.Args) < 3 {
			logo()
			fmt.Printf("not enouth arguments\n")
			os.Exit(1)
		}

		pathDir := os.Args[2]
		logo()
		err := filepath.WalkDir(pathDir, func(path string, file fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if file.IsDir() == false && filepath.Ext(file.Name()) == ".pau" {
				generate_HTML_File(path)
			}

			return nil
		})
		if err != nil {
			logo()
			fmt.Println(err)
			os.Exit(1)
		}

	case "-h", "-help", "help":
		logo()
		fmt.Println("=== help ===")
		fmt.Println("=== run")
		fmt.Println("	run 'pathToFile.pau'")
		fmt.Println("	run .")
		os.Exit(0)
	case  "-v", "-version", "version":
		logo()
		fmt.Printf("pau version %v\n", pauVersion)
		os.Exit(0)
	default:
		logo()
		fmt.Printf("unknown command %q\n", subcomand)
		os.Exit(1)
	}

}

func generate_HTML_File(filepathPau string) {

	content, err := os.ReadFile(filepathPau)
	if err != nil {
		logo()
		fmt.Printf("read file failed! %v\n", err)
		os.Exit(1)
	}
	lexer := syntax.NewLexer(string(content))
	parser := syntax.NewParser(lexer, filepathPau)
	nodes, err := parser.Parse()
	if err != nil {
		logo()
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
		logo()
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
	fmt.Printf("done! %s\n", newfilepathPau)

}

func logo() {
	fmt.Println("#pau")
}