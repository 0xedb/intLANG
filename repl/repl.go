package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/user"

	"github.com/0xedb/intlang/lexer"
	"github.com/0xedb/intlang/token"
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
		for tok := l.NextToken(); tok.Token != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
