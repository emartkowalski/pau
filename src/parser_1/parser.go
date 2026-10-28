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
		// return fmt.Sprintf("%s", n.NodeMainHead.Main_text)
		return n.NodeMainHead.Main_text
	default:
		return ""
	}
}

type NodeKind int

const (
	TypeMainHead NodeKind = iota
	
)

type NodeMainHead struct {
	Main_text string
}

type Node struct {
	NodeType    NodeKind

	NodeMainHead *NodeMainHead

}

// Главная функция парсинга всего файла
func (p *Parser) Parse() ([]*Node, error) {
	var nodes []*Node

	for p.curTok.Type != TokenEOF {

		p.skipNewLine()
		if p.curTok.Type == TokenEOF {
			break
		}

		node, err := p.parseNode()

		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

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
		text := p.curTok.Val
		p.nextToken()

		return newMainHead(text), nil

	default:
		return nil, p.errorf("unknown node type %q", p.curTok.Val)
	}
}

func newMainHead(main_text string) *Node {
	return &Node{
		NodeType: TypeMainHead,
		NodeMainHead: &NodeMainHead{Main_text: main_text}}
}
