package parser

import (
	"fmt"

	"github.com/0xedb/intlang/ast"
	"github.com/0xedb/intlang/lexer"
	"github.com/0xedb/intlang/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	errors    []string
	cur, peek token.TokenObj
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	return p
}

func (p *Parser) nextToken() {
	p.cur = p.peek
	p.peek = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.cur.Token != token.EOF {
		stmt := p.parseStatment()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatment() ast.Statement {
	switch p.cur.Literal {
	case token.AT:
		return p.parseAtStatement()
	case token.RET:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseAtStatement() *ast.AtStatement {
	stmt := &ast.AtStatement{}

	// token
	// identifier
	// value
	stmt.Token = p.cur

	if !p.expectToken(token.IDENT) {
		return nil
	}

	if !p.expectToken(token.ASSIGN) {
		return nil
	}
	stmt.Identifier = &ast.Identifier{
		Token: p.cur,
		Value: p.cur.Literal,
	}

	// skips expressions
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) peekTokenIs(t string) bool {
	return p.peek.Literal == t
}

func (p *Parser) curTokenIs(t string) bool {
	return p.cur.Literal == t
}

func (p *Parser) expectToken(t string) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t string) {
	msg := fmt.Sprintf("expected next token to be %s, but got %s", t, p.peek.Token)

	p.errors = append(p.errors, msg)
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.cur}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
