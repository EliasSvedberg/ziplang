package parser

import (
	"errors"
	"fmt"
	"strconv"
	"ziplang/ast"
	"ziplang/lexer"
	"ziplang/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precendences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.MODULO:   PRODUCT,
	token.LPAREN:   CALL,
}

type Parser struct {
	lexer                *lexer.Lexer
	curToken             token.Token
	peekToken            token.Token
	errors               []string
	prefixParseFunctions map[token.TokenType]func() ast.Expression
	infixParseFunctions  map[token.TokenType]func(ast.Expression) ast.Expression
}

func New(lexer *lexer.Lexer) *Parser {

	p := &Parser{
		lexer: lexer,
	}

	p.prefixParseFunctions = map[token.TokenType]func() ast.Expression{
		token.NUMBER:     p.parseNumberExpression,
		token.IDENTIFIER: p.parseIdentifierExpression,
		token.STRING:     p.parseStringExpression,
		token.MINUS:      p.parsePrefixExpression,
		token.BANG:       p.parsePrefixExpression,
		token.TRUE:       p.parseBooleanExpression,
		token.FALSE:      p.parseBooleanExpression,
		token.LPAREN:     p.parseGroupedExpression,
		token.FUNCTION:   p.parseFunctionExpression,
	}

	p.infixParseFunctions = map[token.TokenType]func(ast.Expression) ast.Expression{
		token.PLUS:     p.parseInfixExpression,
		token.MINUS:    p.parseInfixExpression,
		token.SLASH:    p.parseInfixExpression,
		token.ASTERISK: p.parseInfixExpression,
		token.MODULO:   p.parseInfixExpression,
		token.EQ:       p.parseInfixExpression,
		token.NOT_EQ:   p.parseInfixExpression,
		token.LT:       p.parseInfixExpression,
		token.GT:       p.parseInfixExpression,
		token.LPAREN:   p.parseCallExpression,
	}

	p.advance()
	p.advance()

	return p
}

func (p *Parser) Parse() *ast.Program {

	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !(p.curToken.Type == token.EOF) {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.advance()
	}

	return program
}

func (p *Parser) advance() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {

	switch p.curToken.Type {
	case token.IDENTIFIER:
		return p.parseIdentifierStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() ast.Statement {

	statement := &ast.ExpressionStatement{Token: p.curToken}
	statement.Expression = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.advance()
	}

	return statement
}

func (p *Parser) parseReturnStatement() ast.Statement {

	statement := &ast.ReturnStatement{Token: p.curToken}
	p.advance()

	statement.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.advance()
	}

	return statement
}

func (p *Parser) parseIdentifierStatement() ast.Statement {

	switch p.peekToken.Type {
	case token.CONST:
		return p.parseConstIdentifierStatement()
	case token.VAR:
		return p.parseVarIdentifierStatement()
	case token.ASSIGN:
		return p.parseAssignIdentifierStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseConstIdentifierStatement() ast.Statement {

	statement := &ast.IdentifierStatement{Token: p.curToken}
	p.advance()

	statement.Type = p.curToken

	p.advance()
	statement.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.advance()
	}

	return statement
}

func (p *Parser) parseVarIdentifierStatement() ast.Statement {

	statement := &ast.IdentifierStatement{Token: p.curToken}
	p.advance()

	statement.Type = p.curToken

	p.advance()
	statement.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.advance()
	}

	return statement
}

func (p *Parser) parseAssignIdentifierStatement() ast.Statement {

	statement := &ast.IdentifierStatement{Token: p.curToken}
	p.advance()

	statement.Type = p.curToken

	p.advance()
	statement.Value = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.advance()
	}

	return statement
}

func (p *Parser) parseExpression(precendence int) ast.Expression {
	prefix := p.prefixParseFunctions[p.curToken.Type]

	if prefix == nil {
		msg := fmt.Sprintf("Error: line: %d. Message: no prefix parse function for %s found\n", p.curToken.Line, p.curToken.Type)
		p.errors = append(p.errors, msg)
		return nil
	}

	leftExpression := prefix()

	for !(p.peekToken.Type == token.SEMICOLON) && precendence < p.peekPrecedence() {
		infix := p.infixParseFunctions[p.peekToken.Type]

		if infix == nil {
			return leftExpression
		}

		p.advance()

		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

func (p *Parser) parseInfixExpression(expr ast.Expression) ast.Expression {

	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken,
		Left:     expr,
	}

	precendence := p.currentPrecedence()

	p.advance()
	expression.Right = p.parseExpression(precendence)

	return expression
}

func (p *Parser) parseNumberExpression() ast.Expression {
	number := &ast.NumberExpression{
		Token: p.curToken,
	}

	value, err := strconv.ParseInt(p.curToken.Value, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("Error: line: %d. Message: could not parse %q as integer\n", p.curToken.Line, p.curToken.Value)
		p.errors = append(p.errors, msg)
		return nil
	}

	number.Value = int(value)

	return number
}

func (p *Parser) parseIdentifierExpression() ast.Expression {
	identifier := &ast.IdentifierExpression{
		Token: p.curToken,
		Value: p.curToken.Value,
	}

	return identifier
}

func (p *Parser) parseStringExpression() ast.Expression {
	str := &ast.StringExpression{
		Token: p.curToken,
		Value: p.curToken.Value,
	}

	return str
}

func (p *Parser) parseBooleanExpression() ast.Expression {
	boolean := &ast.BooleanExpression{
		Token: p.curToken,
	}

  if (p.curToken.Value == "true") {
    boolean.Value = true
  } else {
    boolean.Value = false
  }

	return boolean
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken,
	}

	p.advance()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.advance()

	expression := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return expression
}

func (p *Parser) parseFunctionExpression() ast.Expression {
	function := &ast.FunctionExpression{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	function.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	function.Body = p.parseBlockStatement()

	return function
}

func (p *Parser) parseFunctionParameters() []*ast.IdentifierExpression {
	identifiers := []*ast.IdentifierExpression{}

	if p.peekToken.Type == token.RPAREN {
		return identifiers
	}

	p.advance()

	ident := &ast.IdentifierExpression{
		Token: p.curToken,
		Value: p.curToken.Value,
	}

	identifiers = append(identifiers, ident)

	for p.peekToken.Type == token.COMMA {
		p.advance()
		p.advance()
		ident := &ast.IdentifierExpression{
			Token: p.curToken,
			Value: p.curToken.Value,
		}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	blockstatement := &ast.BlockStatement{
		Token: p.curToken,
	}

	blockstatement.Statements = []ast.Statement{}

	p.advance()

	for !(p.curToken.Type == token.RBRACE) && !(p.curToken.Type == token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			blockstatement.Statements = append(blockstatement.Statements, statement)
		}
		p.advance()
	}
	return blockstatement
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	expression := &ast.CallExpression{
		Token:    p.curToken,
		Function: function,
	}

	expression.Arguments = p.parseExpressionList(token.RPAREN, token.COMMA)

	return expression
}

func (p *Parser) parseExpressionList(end token.TokenType, delim token.TokenType) []ast.Expression {
	expressionlist := []ast.Expression{}

	if p.peekToken.Type == end {
		p.advance()
		return expressionlist
	}
	p.advance()

	expressionlist = append(expressionlist, p.parseExpression(LOWEST))

	for p.peekToken.Type == delim {
		p.advance()
		p.advance()

		expressionlist = append(expressionlist, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return expressionlist
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precendences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currentPrecedence() int {
	if p, ok := precendences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.advance()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Error: line: %d. Message: expected next token to be: %s, got %s instead", p.curToken.Line, t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ReportParserErrors() (string, error) {

	errmsg := ""

	if len(p.errors) == 0 {
		return "", nil
	}

	for _, msg := range p.errors {
		errmsg += fmt.Sprintf("parser error: %q", msg)
		errmsg += fmt.Sprintf("Error: line: %d. Message: could not parse %q as integer\n", p.curToken.Line, p.curToken.Value)
	}

	return errmsg, errors.New("Parser reported errors")
}
