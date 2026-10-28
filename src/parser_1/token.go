package parser_1

type TokenType int

type Token struct {
	Type TokenType
	Val  string
	Line int
	Col  int
}

const (
	TokenEOF     TokenType = iota // The end of file
	TokenIllegal                  // Unknown/mistakeble character
	TokenText 					  // text...


	TokenNewLine 				  // '\n'

	TokenOp					      // <
	TokenCl 					  // >
	TokenAnd 				   	  // |

)

// var keywords = map[string]TokenType{

// }

// func LookupIdent(ident string) TokenType {
// 	if tok, ok := keywords[ident]; ok {
// 		return tok
// 	}
// 	return TokenIdent
// }
