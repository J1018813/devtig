package parser

import (
	"github.com/j1018813/devtig/ast"
	"github.com/j1018813/devtig/lexer"
	"github.com/j1018813/devtig/token"
)

// Parser is used to parse tokens into an abstract syntax tree.
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New gives you a new parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.readInitialTokens()
	return p
}

// ParseProgram returns the program which is the root node in the AST.
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

func (p *Parser) readInitialTokens() {
	// Read two tokens, so curToken and peekToken are both set.
	p.nextToken()
	p.nextToken()
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
