package token

type Tok int

const (
	EOF       Tok = iota // end of file
	ILLEGAL              // illegal character
	LET                  // let
	RETURN               // return
	IDENT                // add, foobar, x, y, ...
	INT                  // 123456
	ASSIGN               // =
	PLUS                 // +
	COMMA                // ,
	SEMICOLON            // ;
	LPAREN               // (
	RPAREN               // )
	LBRACE               // {
	RBRACE               // }
)

type Token struct {
	Type    Tok
	Literal string
}

func NewToken(t Tok, ch []byte) Token {
	return Token{Type: t, Literal: string(ch)}
}
