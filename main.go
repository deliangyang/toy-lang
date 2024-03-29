package main

import (
	"os"

	"github.com/deliangyang/tiny-lang/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
