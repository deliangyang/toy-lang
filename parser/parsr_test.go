package parser

import (
	"fmt"
	"testing"

	"github.com/deliangyang/tiny-lang/lexer"
)

func TestParse(t *testing.T) {
	l := lexer.New("let a = 1;")
	p := New(l)
	prog := p.ParseProgram()
	for _, v := range prog.Statements {
		fmt.Println("---->", v.TokenLiteral())
	}
}
