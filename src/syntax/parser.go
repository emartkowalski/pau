package syntax

import (
	"fmt"
	"strings"
)

type Parser struct {
	lexer    *Lexer
	curTok   Token
	filePath string
}

func NewParser(l *Lexer, filepath string) *Parser {
	p := &Parser{lexer: l, filePath: filepath}
	// Read the first token to fill curTok
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.lexer.NextToken()
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
		return strings.ReplaceAll(template, "Title", n.NodeMainHead.Template)	

	case TypeTemplateFloor:
		return n.NodeTemplateFloor.Template

	case TypeHead2:
		return "	" + "<h2>" + n.NodeHead2.TagTitle + "</h2>\n"
	case TypeUl:
		if n.NodeUl.Nodes == nil {
			return "\n"
		}
		var sb strings.Builder

		sb.WriteString("<ul>\n")

		for _, node := range n.NodeUl.Nodes {

			if node.NodeType == TypeA {
				sb.WriteString("	<li>")
				sb.WriteString(strings.ReplaceAll(node.String(), "\n", ""))
				sb.WriteString("</li>")
			}
		}

		sb.WriteString("\n</ul>\n\n")
		return sb.String()

	case TypeTextLine:
		return "	" + "<p>" + n.NodeTextLine.Line + "</p>\n"
	case TypeA:
		return `<a href="` + n.NodeA.Link + `">` + n.NodeA.TitleOfLink + `</a>` + "\n"
	case TypeHr:
		return "	<hr>\n"
	default:
		return ""
	}
}

type NodeKind int
const (
	TypeMainHead NodeKind = iota
	TypeTemplateFloor

	TypeHead2
	TypeUl
	TypeA
	TypeTextLine
	TypeHr
)	

type NodeMainHead struct {
	Template string
}

type NodeTemplateFloor struct {
	Template string
}

type NodeHead2 struct {
	TagTitle string
}

type NodeUl struct {
	Nodes []*Node
}

type NodeA struct {
	Link 	    string
	TitleOfLink string
}

type NodeTextLine struct {
	Line string
}

type NodeHr struct {
	Hr string
}

type Node struct {
	NodeType    NodeKind

	NodeMainHead 	  *NodeMainHead
	NodeTemplateFloor *NodeTemplateFloor

	NodeHead2 		  *NodeHead2
	NodeUl 			  *NodeUl
	NodeA 			  *NodeA
	NodeTextLine 	  *NodeTextLine
	NodeHr 			  *NodeHr

}

// Главная функция парсинга всего файла
func (p *Parser) Parse() ([]*Node, error) {
	var nodes []*Node

	p.skipNewLine()
	if p.curTok.Type == TokenText {
		HeadTitle := p.curTok.Val
		p.nextToken()
		p.skipNewLine()
		nodes = append(nodes, newMainHead(HeadTitle))
	} else {
		return nil, p.errorf("expected for text, got %q", p.curTok.Val)
	}

	p.skipNewLine()
	for p.curTok.Type != TokenEOF {

		node, err := p.parseNode()

		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	nodes = append(nodes, newTamplateFloor())

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
	case TokenTag:

		switch p.curTok.Val {
		case "<h2>":
			return p.parseHead2()
		case "<ul>":
			return p.parseUl()
		case "<a>":
			return p.parseA()
		case "<hr>":
			p.nextToken()
			p.skipNewLine()
			return newHr(), nil
		default:
			return nil, p.errorf("Error parseNode(): there is no tag like this one %q", p.curTok.Val)
		}

	case TokenText:
		return p.parseText()
	default:
		if p.curTok.Type == TokenIllegal {
			return nil, p.errorf("it's not closed , got %q", p.curTok.Val)
		}		
		return nil, p.errorf("unknown node type %q", p.curTok.Val)
	}
}

func (p *Parser) parseHead2() (*Node, error) {
	p.nextToken()

	if p.curTok.Type != TokenText {
		return nil, p.errorf("expected tag, got %q", p.curTok.Val)
	}

	text := p.curTok.Val
	p.nextToken()
	p.skipNewLine()
	
	return newHead2(strings.TrimSpace(text)), nil
}

func (p *Parser) parseUl() (*Node, error) {
	p.nextToken()
	p.skipNewLine()

	Ul := new(NodeUl)

	for p.curTok.Val != "</ul>" || p.curTok.Type != TokenEOF {

		line := p.lexer.CurrentLine(p.curTok.Line)
		for i := 0; i < len(line); i++ {

			switch p.curTok.Type {
			case TokenTag:

				switch p.curTok.Val {
				case "<a>":
					node, err := p.parseA()
					if err != nil {
						return nil, p.errorf("%v", err)
					}

					Ul.Nodes = append(Ul.Nodes, node)
				case "</ul>":
					p.nextToken()
					p.skipNewLine()

					return newUl(Ul.Nodes), nil
				default:
					return nil, p.errorf("there is no tag inside <ul>")
				}
			case TokenText:
				return nil, p.errorf("inside of <ul> you can't use text")
			default:
				return nil, p.errorf("there is no tag like this one %q", p.curTok.Val)
			}
		}
	}

	p.skipNewLine()

	return newUl(Ul.Nodes), nil
}

func (p *Parser) parseA() (*Node, error) {
	p.nextToken()

	A := new(NodeA)

	if p.curTok.Type == TokenString {
		A.Link = p.curTok.Val
		p.nextToken()
	} else {
		return nil, p.errorf("expected link, got %q", p.curTok.Val)
	}

	if p.curTok.Type == TokenAnd {
		p.nextToken()
		// return nil, p.errorf("expected | , got %q", p.curTok.Val)
	}

	if p.curTok.Type == TokenText {
		A.TitleOfLink = p.curTok.Val
		p.nextToken()
	} else if p.curTok.Type == TokenNewLine {
		A.TitleOfLink = A.Link
	} else {
		return nil, p.errorf("expected text for link, got %q", p.curTok.Val)
	}
	
	p.skipNewLine()

	return newA(A.Link, A.TitleOfLink), nil
}

func (p *Parser) parseText() (*Node, error) {

	var sb strings.Builder

	for p.curTok.Type != TokenEOF && p.curTok.Type != TokenTag {
		switch p.curTok.Type {
		case TokenText:
			_, err := sb.WriteString(p.curTok.Val)
			if err != nil {
				return nil, p.errorf("%v", err)
			}
			p.nextToken()

			if p.curTok.Type == TokenNewLine { // skip the first newline

				_, err := sb.WriteString("<br>\n    ")
				if err != nil {
					return nil, p.errorf("%v", err)
				}
				p.nextToken()
			}

			if p.curTok.Type == TokenNewLine { // skip the second newline

				_, err := sb.WriteString("<br>\n    ")
				if err != nil {
					return nil, p.errorf("%v", err)
				}
				p.nextToken()
			}

		case TokenNewLine:
			p.skipNewLine()
		default:
			p.nextToken()
		}
	}

	p.skipNewLine()

	return newTextLine(sb.String()), nil
}

func newMainHead(main_text string) *Node {

	newMainHead := &NodeMainHead{Template: main_text}
	return &Node{
		NodeType: TypeMainHead,
		NodeMainHead: newMainHead}
}

func newTamplateFloor() *Node {
template :=
`
</body>
</html>`
	return &Node{
		NodeType: TypeTemplateFloor,
		NodeTemplateFloor:  &NodeTemplateFloor{template},
	}
}

func newHead2(tagTitle string) *Node {
	return &Node{
		NodeType: TypeHead2,
		NodeHead2: &NodeHead2{TagTitle: tagTitle},
	}
}

func newUl(nodes []*Node) *Node {
	return &Node{
		NodeType: TypeUl,
		NodeUl: &NodeUl{Nodes: nodes},
	}
}

func newA(link string, titleOfLink string) *Node {
	return &Node{
		NodeType: TypeA,
		NodeA: &NodeA{Link: link, TitleOfLink: titleOfLink},
	}
}

func newTextLine(line_text string) *Node {
	return &Node{
		NodeType: TypeTextLine,
		NodeTextLine: &NodeTextLine{Line: line_text},
	}
}

func newHr() *Node {
	return &Node{
		NodeType: TypeHr,
		NodeHr: &NodeHr{},
	}
}
