package parser_1

import (
	"strings"
)

type Lexer struct {
	src  	  string
	pos  	  int // current position in source
	line 	  int // curent line (starts from 1)
	col  	  int // curent column (starts from 1)
	lineStart int
}

func NewLexer(src string) *Lexer {
	return &Lexer{
		src:       src,
		pos:       0,
		line:      1,
		col:  	   1,
		lineStart: 0,
	}
}

func (l *Lexer) CurrentLine(targetLine int) string {
	line := 1
	start := 0
	
	for i := 0; i < len(l.src); i++ {
		if l.src[i] == '\n' {
			if line == targetLine {
				return strings.TrimRight(l.src[start:i], "\r")
			}
			line++
			start = i + 1
		}
	}
	
	// Если это последняя строка файла (без \n в конце)
	if line == targetLine {
		return strings.TrimRight(l.src[start:], "\r")
	}
	
	return ""
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.pos >= len(l.src) {
		return Token{Type: TokenEOF, Line: l.line, Col: l.col}
	}

	ch := l.src[l.pos]
	startCol := l.col

    // if ch == '"' {
    // 	l.advance()
    // 	startPos := l.pos

    // 	for l.pos < len(l.src) && l.src[l.pos] != '"' && l.src[l.pos] != '\n' {
    // 		if l.src[l.pos] == '\\' && l.pos+1 < len(l.src) {
    // 			l.advance()
    // 		}
    // 		l.advance()
    // 	}

    // 	if l.pos >= len(l.src) || l.src[l.pos] != '"' {
    // 		return Token{Type: TokenIllegal, Val: l.src[startPos-1: l.pos], Line: l.line, Col: startCol}
    // 	}

    // 	val := l.src[startPos:l.pos]
    // 	if l.pos < len(l.src) && l.src[l.pos] == '"' {
    // 		l.advance()
    // 	}

    // 	return Token{Type: TokenString, Val: val, Line: l.line, Col: startCol}
    // }


    switch ch {
    case '\n':
    	line := l.line
 		l.advance()
		return Token{Type: TokenNewLine, Val: "\n", Line: line, Col: startCol} 

	case '<':
		l.advance()
		return Token{Type: TokenOp, Val: "<", Line: l.line, Col: startCol}
	case '>':
		l.advance()	
		return Token{Type: TokenCl, Val: ">", Line: l.line, Col: startCol}
	case '|':
		l.advance()
		return Token{Type: TokenAnd, Val: "|", Line: l.line, Col: startCol}
	}

	// handle identifiers
	// Обработка идентификаторов
	startPos := l.pos
	for l.pos < len(l.src) && l.src[l.pos] != '\n' {
		l.advance()
	}

	val := strings.TrimSpace(l.src[startPos:l.pos])
	return Token{
		Type: TokenText,
		Val:  val,
		Line: l.line,
		Col:  startCol,
	}

	// if found unknown character
	// tok := Token{Type: TokenIllegal, Val: string(ch), Line: l.line, Col: startCol}
	// l.advance()
	// return tok
}
// move the position forward by 1 character ahead and track column/line
// Продвигаем указатель на 1 символ вперед и считаем колонку/строку
func (l *Lexer) advance() {
	if l.pos < len(l.src) {
		if l.src[l.pos] == '\n' {
			l.line++
			l.col = 1
			l.lineStart = l.pos + 1
		} else if l.src[l.pos] == '\t' {
			l.col += 4
		} else {
			l.col++
		}
		l.pos++
	}
}
// skips all whitespace, tabs
// Пропуск всех пробелов, табов
func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.src) {

		ch := l.src[l.pos]
		if ch == ' ' || ch == '\t' || ch == '\r' {
			l.advance()
			continue
		}

		// пропускаем коментарии //
		// if ch == '/' && l.pos+1 < len(l.src) && l.src[l.pos+1] == '/' {
		// 	l.advance()
		// 	l.advance()

		// 	for l.pos < len(l.src) && l.src[l.pos] != '\n' {
		// 		l.advance()
		// 	}
		// 	continue
		// }
		break
	}
}

// func (l *Lexer) peek() byte {
// 	if l.pos + 1 >= len(l.src) {
// 		return 0
// 	}
// 	return l.src[l.pos+1]
// }

// func isAlpha(ch byte) bool {
// 	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
// }