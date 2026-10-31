package parser_1

import (
	"fmt"
	"strings"
)

type Parser struct {
	lexer    *Lexer
	curTok   Token
	peekTok  Token
	filePath string
}

func NewParser(l *Lexer, filepath string) *Parser {
	p := &Parser{lexer: l, filePath: filepath}
	// Read the first two tokens to fill curTok, peekTok
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.lexer.NextToken()
}

func (p *Parser) skipNewLine() {
	for p.curTok.Type == TokenNewLine {
		p.nextToken()
	}
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}

	switch n.NodeType {
	case TypeMainHead:
		return n.NodeMainHead.Template
	case TypeTemplateFloor:
		return n.NodeTemplateFloor.Template
	case TypeTextLine:
		return n.NodeTextLine.LineText
	default:
		return ""
	}
}


type NodeKind int
const (
	TypeMainHead NodeKind = iota
	TypeTextLine
	TypeTemplateFloor
)	

type NodeMainHead struct {
	Template string
}

type NodeTemplateFloor struct {
	Template string
}

type NodeTextLine struct {
	LineText string
}


type Node struct {
	NodeType    NodeKind

	NodeMainHead 	  *NodeMainHead
	NodeTextLine 	  *NodeTextLine
	NodeTemplateFloor *NodeTemplateFloor

}

// Главная функция парсинга всего файла
func (p *Parser) Parse() ([]*Node, error) {
	var nodes []*Node
	var HeadTitle string

	if p.curTok.Type == TokenText {
		HeadTitle = p.curTok.Val
		p.nextToken()
		p.skipNewLine()
		nodes = append(nodes, newMainHead(HeadTitle))
	}

	for p.curTok.Type != TokenEOF {

		// p.skipNewLine()
		//if p.curTok.Type == TokenEOF {
		//	break
		//}

		node, err := p.parseNode()

		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	nodes = append(nodes, newTamplateFloor(HeadTitle))

	return nodes, nil
}

func (p *Parser) errorf(format string, args ...any) error {
	msg := fmt.Sprintf(format, args...)
	errLine := fmt.Sprintf("%s", p.lexer.CurrentLine(p.curTok.Line))

	cleanLine := strings.ReplaceAll(errLine, "\t", "    ")

	indent := p.curTok.Col - 1
	if indent < 0 {
		indent = 0
	}
	show := strings.Repeat(" ", indent) + "^"

	lineStr := fmt.Sprintf("%d", p.curTok.Line)
	padding := strings.Repeat(" ", len(lineStr))

	line1 := fmt.Sprintf("%s:%d:%d: %s\n", p.filePath, p.curTok.Line, p.curTok.Col, msg)
	line2 := fmt.Sprintf(" %s | %s\n", lineStr, cleanLine)
	line3 := fmt.Sprintf(" %s | %s", padding, show)
	
	return fmt.Errorf("%v%v%v", line1, line2, line3)
}

func (p *Parser) parseNode() (*Node, error) {

	switch p.curTok.Type {
	case TokenText:
		lineOfText := p.curTok.Val
		p.nextToken()
		p.skipNewLine()
		return newTextLine(lineOfText), nil

	default:
		return nil, p.errorf("unknown node type %q", p.curTok.Val)
	}
}

func newMainHead(main_text string) *Node {
template :=
`<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <style>
        body {
            font-family: Arial, sans-serif;
        }
    </style>
</head>
<body leftmargin="250" rightmargin="250" bgcolor="#d3d3d3">

    <h1 align="center">Title</h1>
    <hr>

`
	newtemplate := strings.ReplaceAll(template, "Title", main_text)

	newMainHead := &NodeMainHead{Template: newtemplate}
	return &Node{
		NodeType: TypeMainHead,
		NodeMainHead: newMainHead}
}

func newTamplateFloor(main_text string) *Node {

template :=
`

    <hr>
    <a href="index.html">HOME</a><br>
    <p>Title | 2026</p>

</body>
</html>
`
	newtemplate := strings.ReplaceAll(template, "Title", main_text)
	nodeTemplateFloor := &NodeTemplateFloor{newtemplate}
	return &Node{
		NodeType: TypeTemplateFloor,
		NodeTemplateFloor: nodeTemplateFloor }
}

func newTextLine(line_text string) *Node {

	nodeTextLine := &NodeTextLine{LineText: line_text}
	return &Node{
		NodeType: TypeTextLine,
		NodeTextLine: nodeTextLine }
}
