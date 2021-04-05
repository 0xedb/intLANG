package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/user"

	"github.com/0xedb/intlang/evaluator"
	"github.com/0xedb/intlang/lexer"
	"github.com/0xedb/intlang/parser"
)

const (
	PROMPT        = ">>> "
	GREET  string = `
	██╗███╗   ██╗████████╗██╗      █████╗ ███╗   ██╗ ██████╗ 
	██║████╗  ██║╚══██╔══╝██║     ██╔══██╗████╗  ██║██╔════╝ 
	██║██╔██╗ ██║   ██║   ██║     ███████║██╔██╗ ██║██║  ███╗
	██║██║╚██╗██║   ██║   ██║     ██╔══██║██║╚██╗██║██║   ██║
	██║██║ ╚████║   ██║   ███████╗██║  ██║██║ ╚████║╚██████╔╝
	╚═╝╚═╝  ╚═══╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ 
																 
	`
)

func StartREPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	fmt.Println(GREET)
	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello, ", user.Name)
	fmt.Println("Welcome to the intLANG programming language")

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, GREET)
	io.WriteString(out, "Woops! We ran into some monkye business here!\n")
	io.WriteString(out, "  parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
