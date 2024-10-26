package lexer

import (
  "testing"
  "ziplang/token"
)


func TestNextToken(t *testing.T) {
  input := "  +  - */=   \n < > \n , ; \n ( )  \n test öra 13 004öra { } \n  [] \n  \"hejhej\" \n ! != == // test-comment \n "

  tests := []struct {
    expectedType token.TokenType
    expectedValue string
    expectedLine int
  }{
    {token.PLUS, "+", 1},
    {token.MINUS, "-", 1},
    {token.ASTERISK, "*", 1},
    {token.SLASH, "/", 1},
    {token.ASSIGN, "=", 1},
    {token.LT, "<", 2},
    {token.GT, ">", 2},
    {token.COMMA, ",", 3},
    {token.SEMICOLON, ";", 3},
    {token.LPAREN, "(", 4},
    {token.RPAREN, ")", 4},
    {token.IDENTIFIER, "test", 5},
    {token.IDENTIFIER, "öra", 5},
    {token.NUMBER, "13", 5},
    {token.NUMBER, "004", 5},
    {token.IDENTIFIER, "öra", 5},
    {token.LBRACE, "{", 5},
    {token.RBRACE, "}", 5},
    {token.LBRACKET, "[", 6},
    {token.RBRACKET, "]", 6},
    {token.STRING, "\"hejhej\"", 7},
    {token.BANG, "!", 8},
    {token.NOT_EQ, "!=", 8},
    {token.EQ, "==", 8},
    {token.COMMENT, "// test-comment ", 8},
    {token.EOF, "EOF", 9},
  }

  l := New(input)

  for i, tt := range tests {
    tok := l.NextToken()

    if tok.Type != tt.expectedType {
      t.Fatalf("tests [%d] - tokentype wrong. expected=%q, got =%q", i, tt.expectedType, tok.Type)
    }

    if tok.Value != tt.expectedValue {
      t.Fatalf("tests [%d] - value wrong. expected=%q, got =%q", i, tt.expectedValue, tok.Value)
    }

    if tok.Line != tt.expectedLine {
      t.Fatalf("tests [%d] - line wrong. expected=%d, got =%d", i, tt.expectedLine, tok.Line)
    }
  }
}
