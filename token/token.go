package token

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
  PLUS      = "+"
  MINUS     = "-"
  ASTERISK  = "*"
  SLASH     = "/"
  ASSIGN    = "="
  BANG      = "!"

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
