package lexer

import (
  "testing"
  "ziplang/token"
)

func TestLexerIllegal(t *testing.T) {
  input := " 3; ${ # ?"

  tests := []struct {
    expectedType token.TokenType
    expectedValue string
    expectedLine int
  }{
    {token.NUMBER, "3", 1},
    {token.SEMICOLON, ";", 1},
    {token.ILLEGAL, "$", 1},
    {token.LBRACE, "{", 1},
    {token.ILLEGAL, "#", 1},
    {token.ILLEGAL, "?", 1},
    {token.EOF, "EOF", 1},
  }

  l := New(input)

  for i, tc := range tests {
    tok := l.NextToken()

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

func TestLexerComment(t *testing.T) {
  input := `
  test := 3; // This is a comment
  // {} +-/""foo bar
  3;
  true; //comment
  `

  tests := []struct {
    expectedType token.TokenType
    expectedValue string
    expectedLine int
  }{
    {token.IDENTIFIER, "test", 2},
    {token.VAR, ":=", 2},
    {token.NUMBER, "3", 2},
    {token.SEMICOLON, ";", 2},
    {token.COMMENT, "// This is a comment", 2},
    {token.COMMENT, `// {} +-/""foo bar`, 3},
    {token.NUMBER, "3", 4},
    {token.SEMICOLON, ";", 4},
    {token.TRUE, "true", 5},
    {token.SEMICOLON, ";", 5},
    {token.COMMENT, "//comment", 5},
    {token.EOF, "EOF", 6},
  }

  l := New(input)

  for i, tc := range tests {
    tok := l.NextToken()

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

func TestLexerKeywords(t *testing.T) {
  input := "true false return fn"

  tests := []struct {
    expectedType token.TokenType
    expectedValue string
    expectedLine int
  }{
    {token.TRUE, "true", 1},
    {token.FALSE, "false", 1},
    {token.RETURN, "return", 1},
    {token.FUNCTION, "fn", 1},
    {token.EOF, "EOF", 1},
  }

  l := New(input)

  for i, tc := range tests {
    tok := l.NextToken()

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

func TestLexerBracket(t *testing.T) {
  input := "[][][[]]"

  tests := []struct {
    expectedType token.TokenType
    expectedValue string
    expectedLine int
  }{
    {token.LBRACKET, "[", 1},
    {token.RBRACKET, "]", 1},
    {token.LBRACKET, "[", 1},
    {token.RBRACKET, "]", 1},
    {token.LBRACKET, "[", 1},
    {token.LBRACKET, "[", 1},
    {token.RBRACKET, "]", 1},
    {token.RBRACKET, "]", 1},
    {token.EOF, "EOF", 1},
  }

  l := New(input)

  for i, tc := range tests {
    tok := l.NextToken()

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

func TestLexerBrace(t *testing.T) {
  input := "{}{}{{}}"

  tests := []struct {
    expectedType token.TokenType
    expectedValue string
    expectedLine int
  }{
    {token.LBRACE, "{", 1},
    {token.RBRACE, "}", 1},
    {token.LBRACE, "{", 1},
    {token.RBRACE, "}", 1},
    {token.LBRACE, "{", 1},
    {token.LBRACE, "{", 1},
    {token.RBRACE, "}", 1},
    {token.RBRACE, "}", 1},
    {token.EOF, "EOF", 1},
  }

  l := New(input)

  for i, tc := range tests {
    tok := l.NextToken()

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

func TestLexerParen(t *testing.T) {
  input := "()(())"

  tests := []struct {
    expectedType token.TokenType
    expectedValue string
    expectedLine int
  }{
    {token.LPAREN, "(", 1},
    {token.RPAREN, ")", 1},
    {token.LPAREN, "(", 1},
    {token.LPAREN, "(", 1},
    {token.RPAREN, ")", 1},
    {token.RPAREN, ")", 1},
    {token.EOF, "EOF", 1},
  }

  l := New(input)

  for i, tc := range tests {
    tok := l.NextToken()

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

func TestLexerDelimiters(t *testing.T) {
  input := ",;,,;;"

  tests := []struct {
    expectedType token.TokenType
    expectedValue string
    expectedLine int
  }{
    {token.COMMA, ",", 1},
    {token.SEMICOLON, ";", 1},
    {token.COMMA, ",", 1},
    {token.COMMA, ",", 1},
    {token.SEMICOLON, ";", 1},
    {token.SEMICOLON, ";", 1},
    {token.EOF, "EOF", 1},
  }

  l := New(input)

  for i, tc := range tests {
    tok := l.NextToken()

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

func TestLexerOperators(t *testing.T) {
  input := "+-*%/=!:::=<>==!="

  tests := []struct {
    expectedType token.TokenType
    expectedValue string
    expectedLine int
  }{
    {token.PLUS, "+", 1},
    {token.MINUS, "-", 1},
    {token.ASTERISK, "*", 1},
    {token.MODULO, "%", 1},
    {token.SLASH, "/", 1},
    {token.ASSIGN, "=", 1},
    {token.BANG, "!", 1},
    {token.CONST, "::", 1},
    {token.VAR, ":=", 1},
    {token.LT, "<", 1},
    {token.GT, ">", 1},
    {token.EQ, "==", 1},
    {token.NOT_EQ, "!=", 1},
    {token.EOF, "EOF", 1},
  }

  l := New(input)

  for i, tc := range tests {
    tok := l.NextToken()

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



func TestNextTokenGeneric(t *testing.T) {
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
    tok := l.NextToken()

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
