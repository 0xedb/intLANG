package ast

import (
	"github.com/0xedb/intlang/token"
)

type Node interface {
	TokenValue() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.TokenObj // IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenValue() string {
	return i.Token.Literal
}

type AtStatement struct {
	Token      token.TokenObj
	Identifier *Identifier
	Value      Expression
}

func (a *AtStatement) statementNode() {}
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

func (r *ReturnStatement) statementNode() {}

func (r *ReturnStatement) TokenValue() string {
	return r.Token.Literal
}

type ExpressionStatement struct {
	Token      token.TokenObj
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}

func (e *ExpressionStatement) TokenValue() string {
	return e.Token.Literal
}

type IntegralExpression struct {
	Token token.TokenObj
	Value int64
}

func (i *IntegralExpression) expressionNode() {}
func (i *IntegralExpression) TokenValue() string {
	return i.Token.Literal
}

type PrefixExpression struct {
	Token    token.TokenObj
	Operator string
	Right    Expression
}

func (p *PrefixExpression) expressionNode() {}
func (p *PrefixExpression) TokenValue() string {
	return p.Token.Literal
}

type InfixExpression struct {
	Token       token.TokenObj
	Operator    string
	Left, Right Expression
}

func (i *InfixExpression) expressionNode() {}
func (i *InfixExpression) TokenValue() string {
	return i.Token.Literal
}

type Boolean struct {
	Token token.TokenObj
	Value bool
}

func (b *Boolean) expressionNode()    {}
func (b *Boolean) TokenValue() string { return b.Token.Literal }

type IfExpression struct {
	Token       token.TokenObj // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()    {}
func (ie *IfExpression) TokenValue() string { return ie.Token.Literal }

type BlockStatement struct {
	Token      token.TokenObj // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()     {}
func (bs *BlockStatement) TokenValue() string { return bs.Token.Literal }

type FunctionLiteral struct {
	Token      token.TokenObj // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()    {}
func (fl *FunctionLiteral) TokenValue() string { return fl.Token.Literal }

type CallExpression struct {
	Token     token.TokenObj // The '(' token
	Function  Expression     // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()    {}
func (ce *CallExpression) TokenValue() string { return ce.Token.Literal }
