package parser

import (
	"Writing-Interpreter-In-Go/ast"
	"Writing-Interpreter-In-Go/lexer"
	"Writing-Interpreter-In-Go/token"
)

// Parser ...
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New return a pointer to the parser with a given lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram ...
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
