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
		fmt.Printf("run 'example.p'")
		os.Exit(0)
	default:
		fmt.Println("no more commands only run")
		os.Exit(1)
	}

}

func doWork(filepath string) {
	tamplate1 :=
`<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>Personal Wiki</title>
    <style>
        body {
            font-family: Arial, sans-serif;
        }
    </style>
</head>
<body leftmargin="250" rightmargin="250" bgcolor="#d3d3d3">

    <h1 align="center">Personal Wiki</h1>
    <hr>
    
`

tamplate2 :=
`
    <hr>
    <a href="index.html">HOME</a><br>
    <p>Personal Wiki | 2026</p>

</body>
</html>
`

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

	fmt.Print(tamplate1)
	for _, node := range nodes {
		fmt.Print(node.String())
	}
	fmt.Print(tamplate2)
}

/*
`

*/


/*
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>Personal Wiki</title>
    <style>
        body {
            font-family: Arial, sans-serif;
        }
    </style>
</head>
<body leftmargin="250" rightmargin="250" bgcolor="#d3d3d3">

    <h1 align="center">Personal Wiki</h1>
    <hr>

    <hr>
    <a href="index.html">HOME</a><br>
    <p>Personal Wiki | 2026</p>

</body>
</html>
*/
