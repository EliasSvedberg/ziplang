package token

import (
  "bytes"
  "strconv"
)

type Token struct {
  Type   TokenType
  Value  string
  Line   int
}

func New(t TokenType, v string, l int) Token {
  return Token {
    Type:   t,
    Value:  v,
    Line:   l,
  }
}

func (t *Token) ToString() string {
  var out bytes.Buffer

  out.WriteString("Token {\n")
  out.WriteString("Type: ")
  out.WriteString(string(t.Type))
  out.WriteString(",\n")
  out.WriteString("Value: ")
  out.WriteString(t.Value)
  out.WriteString(",\n")
  out.WriteString("Line: ")
  out.WriteString(strconv.Itoa(t.Line))
  out.WriteString(",\n")
  out.WriteString("}")

  return out.String()
}

type TokenType string

const (
  ILLEGAL  = "ILLEGAL"
  EOF      = "EOF"
  COMMENT  = "COMMENT"

  // Identifiers + literals
  IDENTIFIER = "IDENTIFIER"
  NUMBER     = "NUMBER"
  STRING     = "STRING"

  // Operators
  PLUS      = "PLUS"
  MINUS     = "MINUS"
  ASTERISK  = "ASTERISK"
  SLASH     = "SLASH"
  ASSIGN    = "ASSIGN"
  BANG      = "BANG"

  // Comparisons
  LT = "<"
  GT = ">"
  EQ = "=="
  NOT_EQ = "!="

  // Delimiters
  COMMA = ","
  SEMICOLON = ";"
  //COLON = ":"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"
  LBRACKET = "["
  RBRACKET = "]"

)
