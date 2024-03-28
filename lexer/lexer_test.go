package lexer

import (
	"testing"

	"github.com/deliangyang/tiny-lang/token"
)

func TestLexer(t *testing.T) {
	type Token struct {
		Type    token.Tok
		Literal string
	}
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			`=+(){},;`,
			[]Token{
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}
	for _, test := range tests {
		l := New(test.input)
		for _, expected := range test.expected {
			tok := l.NextToken()
			if tok.Type != expected.Type {
				t.Fatalf("Expected token type %q, got %q", expected.Type, tok.Type)
			}
			if tok.Literal != expected.Literal {
				t.Fatalf("Expected literal %q, got %q", expected.Literal, tok.Literal)
			}
		}
	}
}
