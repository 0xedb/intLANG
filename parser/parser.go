package parser

import (
	"fmt"
	"strconv"

	"github.com/0xedb/intlang/ast"
	"github.com/0xedb/intlang/lexer"
	"github.com/0xedb/intlang/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer     *lexer.Lexer
	errors    []string
	cur, peek token.TokenObj

	infixFn  map[token.Token]infixParseFn
	prefixFn map[token.Token]prefixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:    l,
		errors:   []string{},
		infixFn:  map[token.Token]infixParseFn{},
		prefixFn: map[token.Token]prefixParseFn{},
	}

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegralLiteral)
	p.registerPrefix(token.NOT, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.DIV, p.parseInfixExpression)
	p.registerInfix(token.MULT, p.parseInfixExpression)
	p.registerInfix(token.EQL, p.parseInfixExpression)
	p.registerInfix(token.NEQL, p.parseInfixExpression)
	p.registerInfix(token.LST, p.parseInfixExpression)
	p.registerInfix(token.GRT, p.parseInfixExpression)

	return p
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.cur}
	if !p.expectToken(token.LPAREN) {
		return nil
	}
	p.nextToken()
	expression.Condition = p.parseExpression(token.LOWEST)
	if !p.expectToken(token.RPAREN) {
		return nil
	}
	if !p.expectToken(token.LCURL) {
		return nil
	}
	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.EL) {
		p.nextToken()
		if !p.expectToken(token.LCURL) {
			return nil
		}
		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.cur}
	block.Statements = []ast.Statement{}
	p.nextToken()
	for !p.curTokenIs(token.RCURL) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(token.LOWEST)

	if !p.expectToken(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.cur, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.cur,
		Operator: p.cur.Literal,
		Left:     left,
	}

	precedence := token.LookupPrecedence(p.cur.Literal)

	p.nextToken()

	exp.Right = p.parseExpression(precedence)

	return exp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.cur,
		Operator: p.cur.Literal,
	}

	p.nextToken()
	exp.Right = p.parseExpression(token.PREFIX)

	return exp
}

func (p *Parser) noPrefixParseFnError(t string) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseIntegralLiteral() ast.Expression {
	lit := &ast.IntegralExpression{Token: p.cur}

	value, err := strconv.ParseInt(p.cur.Literal, 0, 64)

	if err == nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.cur.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parseIdentifier() ast.Expression {

	return &ast.Identifier{Token: p.cur, Value: p.cur.Literal}
}

func (p *Parser) registerInfix(tok token.Token, fn infixParseFn) {
	p.infixFn[tok] = fn
}

func (p *Parser) registerPrefix(tok token.Token, fn prefixParseFn) {
	p.prefixFn[tok] = fn
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
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.cur.Literal {
	case token.AT:
		return p.parseAtStatement()
	case token.RET:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.cur}

	stmt.Expression = p.parseExpression(token.LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixFn[p.cur.Token]

	if prefix == nil {
		p.noPrefixParseFnError(p.cur.Literal)
		return nil
	}

	leftExp := prefix()

	for p.cur.Token != token.SEMICOLON && precedence < token.LookupPrecedence(p.peek.Literal) {
		infix := p.infixFn[p.cur.Token]

		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
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
