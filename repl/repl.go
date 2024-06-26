package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/deliangyang/tiny-lang/lexer"
	"github.com/deliangyang/tiny-lang/token"
)

const PORMPT = ">>>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PORMPT)

		scan := scanner.Scan()
		if !scan {
			return
		}
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
