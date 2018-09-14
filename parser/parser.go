package parser

import (
	"chango/ast"
	"chango/lexer"
	"chango/token"
)

// Parser ...
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token // Current token under examination
	peekToken token.Token // what to do next
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// New ...
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// ParseProgram ...
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
