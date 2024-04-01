package parser

import (
	"fmt"
	"testing"

	"github.com/deliangyang/tiny-lang/lexer"
)

func TestParse(t *testing.T) {
	l := lexer.New(`
let x = if (4 > 5) { 4 } else { 5 };
`)
	p := New(l)
	prog := p.ParseProgram()
	for _, v := range prog.Statements {
		fmt.Println("---->", v)
	}
}
