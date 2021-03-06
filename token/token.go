package token

import (
	"regexp"
)

type (
	Token string

	TokenObj struct {
		Token   Token
		Literal string
	}

	none struct{}
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	COMMENT = "COMMENT"

	INT    = "INT"
	IDENT  = "IDENT"
	STRING = "\""

	PLUS  = "+"
	MINUS = "-"
	MULT  = "*"
	DIV   = "/"
	NOT   = "!"

	ASSIGN = "="
	EQL    = "=="
	NEQL   = "!="
	LST    = "<"
	GRT    = ">"

	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	RPAREN    = ")"
	LPAREN    = "("
	RCURL     = "}"
	LCURL     = "{"
	RBRAC     = "]"
	LBRAC     = "["

	FUNCTION = "fn"
	AT       = "@"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	EL       = "el"
	RET      = "ret"
)

var keywords map[string]none
var precedence map[string]int

func init() {
	keywords = map[string]none{
		FUNCTION: none{},
		AT:       none{},
		TRUE:     none{},
		FALSE:    none{},
		IF:       none{},
		EL:       none{},
		RET:      none{},
	}

	precedence = map[string]int{
		ASSIGN: EQUALS,
		NEQL:   EQUALS,
		LST:    LESSGREATER,
		GRT:    LESSGREATER,
		PLUS:   SUM,
		MINUS:  SUM,
		DIV:    PRODUCT,
		MULT:   PRODUCT,
	}
}

func LookupIdentifier(id string) Token {
	if _, found := keywords[id]; found {
		return Token(id)
	}
	return IDENT
}

func (tok Token) IsKeyWord() bool {
	_, found := keywords[string(tok)]
	return found
}

func IsLetter(tok byte) bool {
	re := regexp.MustCompile(`[a-zA-Z_]`)

	return re.MatchString(string(tok))
}

func IsNumber(tok byte) bool {
	re := regexp.MustCompile(`[0-9]`)

	return re.MatchString(string(tok))
}

func LookupPrecedence(tok string) int {
	if val, ok := precedence[tok]; ok {
		return val
	}

	return LOWEST
}
