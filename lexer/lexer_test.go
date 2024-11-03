package lexer

import (
  "testing"
  "ziplang/token"
)


func TestNextToken(t *testing.T) {
  input := "  +  - */=   \n < > \n , ; \n ( )  \n test öra 13 004öra { } \n  [] \n  \"hejhej\" \n ! != == // test-comment \n :: :=  % false true return fn"

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
    {token.CONST, "::", 9},
    {token.VAR, ":=", 9},
    {token.MODULO, "%", 9},
    {token.FALSE, "false", 9},
    {token.TRUE, "true", 9},
    {token.RETURN, "return", 9},
    {token.FUNCTION, "fn", 9},
    {token.EOF, "EOF", 9},
  }

  l := New(input)

  for i, tc := range tests {
    tok := l.Nextcoken()

    if tok.Type != tc.expectedType {
      t.Fatalf("tests [%d] - tokentype wrong. expected=%q, got =%q", i, tc.expectedType, tok.Type)
    }

    if tok.Value != tc.expectedValue {
      t.Fatalf("tests [%d] - value wrong. expected=%q, got =%q", i, tc.expectedValue, tok.Value)
    }

    if tok.Line != tc.expectedLine {
      t.Fatalf("tests [%d] - line wrong. expected=%d, got =%d", i, tc.expectedLine, tok.Line)
    }
  }
}
