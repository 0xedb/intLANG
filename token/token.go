package token

type (
	Token uint8

	TokenObj struct {
		Token   Token
		Literal string
	}
)

const (
	ILLEGAL Token = iota
	EOF
	COMMENT

	// literal
	INT
	IDENT
	STRING

	// operator
	PLUS
	MINUS
	MULT
	DIV
	NOT

	ASSIGN
	EQL
	NEQL
	LST
	GRT

	// delimiter
	COMMA
	COLON
	SEMICOLON
	RPAREN
	LPAREN
	RCURL
	LCURL
	RBRAC
	LBRAC

	// keyword
	keyword_beg
	FUNCTION
	AT
	TRUE
	FALSE
	IF
	EL
	RET
	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	INT:    "INT",
	IDENT:  "IDENT",
	STRING: "\"",

	PLUS:  "+",
	MINUS: "-",
	MULT:  "*",
	DIV:   "/",
	NOT:   "!",

	ASSIGN: ":=",
	EQL:  "==",
	NEQL: "!=",
	LST:  "<",
	GRT:  ">",

	COMMA:     ",",
	COLON:     ":",
	SEMICOLON: ";",
	RPAREN:    ")",
	LPAREN:    "(",
	RCURL:     "}",
	LCURL:     "{",
	RBRAC:     "]",
	LBRAC:     "[",

	FUNCTION: "func",
	AT:       "@",
	TRUE:     "true",
	FALSE:    "false",
	IF:       "if",
	EL:       "el",
	RET:      "ret",
}

var keywords map[string]Token

func init() {
	keywords = map[string]Token{}

	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

func LookupIdentifier(id string) Token {
	if tok, found := keywords[id]; found {
		return tok
	}

	return IDENT
}

func (tok Token) IsKeyWord() bool {
	return keyword_beg < tok && tok < keyword_end
}
