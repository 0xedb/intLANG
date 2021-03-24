package lexer

import (
	"github.com/0xedb/intlang/token"
)

type Lexer struct {
	input       string
	pos, offset int
	ch          byte
}

func New(input string) *Lexer {
	l := new(Lexer)
	l.input = input

	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.offset >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.offset]
	}

	l.pos = l.offset
	l.offset++
}

func (l *Lexer) NextToken() token.TokenObj {
	var tok token.TokenObj

	l.eatWhitespace()

	switch string(l.ch) {
	case token.PLUS:
		tok = makeToken(token.PLUS, l.ch)
	case token.MINUS:
		tok = makeToken(token.MINUS, l.ch)
	case token.MULT:
		tok = makeToken(token.MULT, l.ch)
	case token.DIV:
		tok = makeToken(token.DIV, l.ch)
	case token.NOT:
		tok = makeToken(token.NOT, l.ch)
	case token.EQL:
		// assignment operator here too
		tok = makeToken(token.EQL, l.ch)
	case token.NEQL:
		tok = makeToken(token.NEQL, l.ch)
	case token.LST:
		tok = makeToken(token.LST, l.ch)
	case token.GRT:
		tok = makeToken(token.GRT, l.ch)
	case token.COMMA:
		tok = makeToken(token.COMMA, l.ch)
	case token.COLON:
		tok = makeToken(token.COLON, l.ch)
	case token.SEMICOLON:
		tok = makeToken(token.SEMICOLON, l.ch)
	case token.RPAREN:
		tok = makeToken(token.RPAREN, l.ch)
	case token.LPAREN:
		tok = makeToken(token.LPAREN, l.ch)
	case token.RCURL:
		tok = makeToken(token.RCURL, l.ch)
	case token.LCURL:
		tok = makeToken(token.LCURL, l.ch)
	case token.RBRAC:
		tok = makeToken(token.RBRAC, l.ch)
	case token.LBRAC:
		tok = makeToken(token.LBRAC, l.ch)
	case token.STRING:
		tok = makeToken(token.STRING, l.ch)
	case token.AT:
		tok = makeToken(token.AT, l.ch)
	case token.ASSIGN:
		tok = makeToken(token.ASSIGN, l.ch)
	case string(byte(0)):
		tok = makeToken(token.EOF, l.ch)
	default:
		if token.IsNumber(l.ch) {
			tok.Token = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else if token.IsLetter(l.ch) {
			tok.Token = token.IDENT
			tok.Literal = l.readIdentifier()
			return tok
		} else {
			tok = makeToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

func makeToken(tok token.Token, ch byte) token.TokenObj {
	return token.TokenObj{Token: tok, Literal: string(ch)}
}

func (l *Lexer) readNumber() string {
	cur := l.pos

	for token.IsNumber(l.ch) {
		l.readChar()
	}

	return l.input[cur:l.pos]
}

func (l *Lexer) readIdentifier() string {
	cur := l.pos

	for token.IsLetter(l.ch) {
		l.readChar()
	}

	return l.input[cur:l.pos]
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
