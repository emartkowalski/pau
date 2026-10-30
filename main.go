package main

import (
	"fmt"
	"pau/src/parser_1"
	"os"
)

func main() {

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

	filePath := "main.pau"
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
