package ast

import (
	"github.com/0xedb/intlang/token"
)

type Node interface {
	TokenValue() string
}

type Statement interface {
	Node
	statmentNode()
}

type Expression interface {
	Node
	expressionNode()
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

func (a *AtStatement) statmentNode() {}
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

type ReturnStatement struct {
	Token       token.TokenObj
	ReturnValue Expression
}

func (r *ReturnStatement) statmentNode() {}

func (r *ReturnStatement) TokenValue() string {
	return r.Token.Literal
}

type ExpressionStatement struct {
	Token      token.TokenObj
	Expression Expression
}

func (e *ExpressionStatement) statmentNode() {}

func (e *ExpressionStatement) TokenValue() string {
	return e.Token.Literal
}
