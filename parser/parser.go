package parser

import (
	"ziplang/lexer"
	"ziplang/token"
	"ziplang/ast"
  "strconv"
  "fmt"
  "errors"
)
const (
  _ int = iota
  LOWEST
  SUM
  PRODUCT
)

var precendences = map[token.TokenType]int{
  token.PLUS: SUM,
  token.MINUS: SUM,
  token.SLASH: PRODUCT,
  token.ASTERISK: PRODUCT,
}

type Parser struct {
	lexer *lexer.Lexer
	curToken token.Token
  peekToken token.Token
  errors []string
  prefixParseFunctions map[token.TokenType]func() ast.Expression
  infixParseFunctions map[token.TokenType]func(ast.Expression) ast.Expression
}

func New(lexer *lexer.Lexer) *Parser {

  p := &Parser {
    lexer: lexer,
  }

  p.prefixParseFunctions =  map[token.TokenType]func() ast.Expression {
    token.NUMBER: p.parseNumberExpression,
  }

  p.infixParseFunctions =  map[token.TokenType]func(ast.Expression) ast.Expression {
    token.PLUS: p.parseInfixExpression,
    token.MINUS: p.parseInfixExpression,
    token.SLASH: p.parseInfixExpression,
    token.ASTERISK: p.parseInfixExpression,
  }

  p.advance()
  p.advance()

  return p
}

func (p *Parser) Parse() *ast.Program {
  program := &ast.Program{}
  program.Expressions = []ast.Expression{}

  for !(p.curToken.Type == token.EOF) {
    expression := p.parseExpression(LOWEST)

    if p.peekToken.Type == token.SEMICOLON {
      p.advance()
    }

    if expression != nil {
      program.Expressions = append(program.Expressions, expression)
    }

    p.advance()
  }

  return program
}

func (p *Parser) advance() {
  p.curToken = p.peekToken
  p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseExpression(precendence int) ast.Expression {
  prefix := p.prefixParseFunctions[p.curToken.Type]

  if prefix == nil {
    msg := fmt.Sprintf("Error: line: %d. Message: no prefix parse function for %s found", p.curToken.Line, p.curToken.Type)
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
    Token: p.curToken,
    Operator: p.curToken,
    Left: expr,
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
    msg := fmt.Sprintf("Error: line: %d. Message: could not parse %q as integer", p.curToken.Line, p.curToken.Value)
    p.errors = append(p.errors, msg)
    return nil
  }

  number.Value = int(value)

  return number
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

func (p *Parser) ReportParserErrors() (string, error) {

  errmsg := ""
	
	if len(p.errors) == 0 {
		return "", nil
	}

	for _, msg := range p.errors {
		errmsg += fmt.Sprintf("parser error: %q", msg)
    errmsg += fmt.Sprintf("Error: line: %d. Message: could not parse %q as integer", p.curToken.Line, p.curToken.Value)
	}

  return errmsg, errors.New("Parser reported errors")
}
