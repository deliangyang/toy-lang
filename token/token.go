package token

type Tok int

var tokStrings = [...]string{
	EOF:       "EOF",
	ILLEGAL:   "ILLEGAL",
	LET:       "LET",
	RETURN:    "RETURN",
	IDENT:     "IDENT",
	INT:       "INT",
	ASSIGN:    "=",
	PLUS:      "+",
	COMMA:     ",",
	SEMICOLON: ";",
	LPAREN:    "(",
	RPAREN:    ")",
	LBRACE:    "{",
	RBRACE:    "}",
	EQ:        "==",
	NOT_EQ:    "!=",
	GT:        ">",
	LT:        "<",
	GTE:       ">=",
	LTE:       "<=",
	FUNCTION:  "FN",
	BANG:      "!",
	MINUS:     "-",
	ASTERISK:  "*",
	SLASH:     "/",
	IF:        "IF",
	ELSE:      "ELSE",
}

func (t Tok) String() string {
	return tokStrings[t]
}

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
	EQ                   // ==
	NOT_EQ               // !=
	GT                   // >
	LT                   // <
	GTE                  // >=
	LTE                  // <=
	FUNCTION             // fn
	BANG                 // !
	MINUS                // -
	ASTERISK             // *
	SLASH                // /
	IF                   // if
	ELSE                 // else
)

type Token struct {
	Type    Tok
	Literal string
}

func NewToken(t Tok, ch []byte) Token {
	return Token{Type: t, Literal: string(ch)}
}

var keywords = map[string]Tok{
	"let":    LET,
	"return": RETURN,
	"fn":     FUNCTION,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) Tok {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
