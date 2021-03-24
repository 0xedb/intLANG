package lexer

import (
	"fmt"
	"testing"

	"github.com/0xedb/intlang/token"
)

func TestNextToken(t *testing.T) {
	input := `
	@ five = 5;
	@ ten = 10;
	@ add = fn(x, y) {
	x + y;
	};
	@ result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	if (5 < 10) {
	return true;
	} else {
	return false;
	}
	10 == 10;
	10 != 9;
	`

	lex := New(input)

	for output := lex.NextToken(); output.Token != token.EOF; output = lex.NextToken() {
		fmt.Println(output)
	}

	t.Log("done")
}
