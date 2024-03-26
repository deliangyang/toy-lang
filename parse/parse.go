package parse

import (
	"github.com/deliangyang/tiny-lang/ast"
	"github.com/deliangyang/tiny-lang/lexer"
	"github.com/deliangyang/tiny-lang/token"
)

const (
	_ int = iota
	// LOWEST is the lowest precedence
	LOWEST
	// EQUALS is the equals precedence
	EQUALS
	// LESSGREATER is the less greater precedence
	LESSGREATER
	// SUM is the sum precedence
	SUM
	// PRODUCT is the product precedence
	PRODUCT
	// PREFIX is the prefix precedence
	PREFIX
	// CALL is the call precedence
	CALL
)

// Parser is the struct for the parser
type Parser struct {
	l              *lexer.Lexer
	errors         []string
	curToken       token.Token
	peekToken      token.Token
	prefixParseFns map[token.Tok]prefixParseFn
	infixParseeFns map[token.Tok]infixParseFn
}

type (
	prefixParseFn func() ast.Statement
	infixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(t token.Tok, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

func (p *Parser) registerInfix(t token.Tok, fn infixParseFn) {
	p.infixParseeFns[t] = fn
}

// New creates a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

// Errors returns the parser's errors
func (p *Parser) Errors() []string {
	return p.errors
}

// nextToken sets the current token to the next token and the peek token to the next next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseStatement parses a statement
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseLetStatement parses a let statement
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	return stmt
}

// expectPeek checks if the next token is of the expected type
func (p *Parser) expectPeek(t token.Tok) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// peekError appends an error to the parser's errors
func (p *Parser) peekError(t token.Tok) {
	msg := "expected next token to be %s, got %s instead"
	p.errors = append(p.errors, msg)
}

// curTokenIs checks if the current token is of the expected type
func (p *Parser) curTokenIs(t token.Tok) bool {
	return p.curToken.Type == t
}

// peekTokenIs checks if the peek token is of the expected type
func (p *Parser) peekTokenIs(t token.Tok) bool {
	return p.peekToken.Type == t
}

// parseReturnStatement parses a return statement
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	return stmt
}

// parseExpressionStatement parses an expression statement
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression()

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpression parses an expression
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()
	return leftExp
}
