package syntax

import (
	"testing"
	"strings"
)



func TestParseSrcTrue(t *testing.T) {
	table_test := []struct{
		name    string
		input   string
		expect  string
	}{
		{
			name: "test1",
			input:
`Personal Wiki

<h2> вступление
<ul>
<li> асталависта бейбе
</ul>

hello world
another text

the end of text...

<hr>
<a> "index.html" | HOME
Main Site | 2026
`,
			expect:
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

	<h2>вступление</h2>
<ul>
	<li>асталависта бейбе</li>
</ul>
	<p>hello world<br>
    another text<br>
    <br>
    the end of text...<br>
    <br>
    </p>
	<hr>
<a href="index.html">HOME</a>
	<p>Main Site | 2026<br>
    </p>

</body>
</html>`,
	},
		{
			name: "test_1.1",
			input:
`Personal Wiki

<h2> вступление
<ul>
<li> асталависта бейбе
</ul>

hello world
another text


the end of text...

<hr>
<a> "index.html" | HOME
Main Site | 2026
`,
			expect:
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

	<h2>вступление</h2>
<ul>
	<li>асталависта бейбе</li>
</ul>
	<p>hello world<br>
    another text<br>
    <br>
    the end of text...<br>
    <br>
    </p>
	<hr>
<a href="index.html">HOME</a>
	<p>Main Site | 2026<br>
    </p>

</body>
</html>`,
		},
}
	for _, tt := range table_test {
		t.Run(tt.input, func(t *testing.T) {
			var sb strings.Builder

			lexer := NewLexer(tt.input)
			parser := NewParser(lexer, "test.lt")

			nodes, err := parser.Parse()
			if err != nil {
				t.Fatalf("unexpected parse error: %v", err)
			}

			if len(nodes) == 0 {
				t.Fatalf("expected 1 node, got 0")
			}

			for _, node := range nodes {
				_, err := sb.WriteString(node.String())
				if err != nil {
					t.Fatalf("%q\n", err)
				}
			}

			got := sb.String()

			if got != tt.expect {
				t.Fatalf("\n======== input ========\n%s\n======== expect ========\n%s\n======== got ========\n%s",
					tt.input, tt.expect, got)
			}
		})
	}
}
