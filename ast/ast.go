package ast

import (
	"github.com/0xedb/intlang/token"
)

type Node interface {
	TokenValue() string
}

type Statement struct {
	Node
}

type Expression struct {
	Node
}

type Identifier struct {
	Token token.TokenObj // IDENT
	Value string
}

func (i *Identifier) TokenValue() string {
	return i.Token.Literal
}

type AtStatement struct {
	Token      token.TokenObj
	Identifier *Identifier
	Value      Expression
}

func (a *AtStatement) TokenValue() string {
	return a.Token.Literal
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenValue() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenValue()
	}

	return ""
}
