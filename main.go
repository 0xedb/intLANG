package main

import (
	"os"

	"github.com/0xedb/intlang/repl"
)

func main() {
	repl.StartREPL(os.Stdout, os.Stdout)
}
