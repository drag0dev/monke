package parser

import (
    "monke/ast"
    "monke/lexer"
    "monke/token"
)

type Parser struct {
    l *lexer.Lexer
    currToken token.Token
    peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}

    // read twice so currToken and peekToken are set
    p.NextToken()
    p.NextToken()

    return p
}

func (p *Parser) NextToken() {
    p.currToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    return nil
}
