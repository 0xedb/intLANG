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

	FUNCTION = "func"
	AT       = "@"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	EL       = "el"
	RET      = "ret"
)

var keywords map[string]none

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
