package lexer

import (
  "ziplang/token"
  "unicode"
  "unicode/utf8"
)

type Lexer struct {
  source   string
  char     rune
  position int
  line     int
}

func New(source string) *Lexer {
  l := &Lexer{
    source:   source,
    char:     0,
    position: 0,
    line:     1,
  }

  return l
}

func (lexer *Lexer) NextToken() token.Token {

  //read char and skip whitespace
  lexer.readChar()
  lexer.skipWhiteSpace()

  //switch on char and return token

  switch lexer.char {
  case 0:
    return token.New(token.EOF, "EOF", lexer.line)
  case '+':
    return token.New(token.PLUS, string(lexer.char), lexer.line)
  case '-':
    return token.New(token.MINUS, string(lexer.char), lexer.line)
  case '*':
    return token.New(token.ASTERISK, string(lexer.char), lexer.line)
  case '%':
    return token.New(token.MODULO, string(lexer.char), lexer.line)

  // SLASH, COMMENT
  case '/':
    if lexer.peekChar() == '/' {
      return lexer.readComment()
    }
    return token.New(token.SLASH, string(lexer.char), lexer.line)
  // Assign, EQ
  case '=':
    if lexer.peekChar() == '=' {
      lexer.readChar()
      return token.New(token.EQ, string("=="), lexer.line)
    }
    return token.New(token.ASSIGN, string(lexer.char), lexer.line)

  // BANG, NOT EQ
  case '!':
    if lexer.peekChar() == '=' {
      lexer.readChar()
      return token.New(token.NOT_EQ, string("!="), lexer.line)
    }
    return token.New(token.BANG, string(lexer.char), lexer.line)

  // VAR, CONST
  case ':':
    if lexer.peekChar() == ':' {
      lexer.readChar()
      return token.New(token.CONST, string("::"), lexer.line)
    } else if lexer.peekChar() == '=' {
      lexer.readChar()
      return token.New(token.VAR, string(":="), lexer.line)
    }
    return token.New(token.ILLEGAL, string(lexer.char), lexer.line)

  case '<':
    return token.New(token.LT, string(lexer.char), lexer.line)
  case '>':
    return token.New(token.GT, string(lexer.char), lexer.line)
  case ',':
    return token.New(token.COMMA, string(lexer.char), lexer.line)
  case ';':
    return token.New(token.SEMICOLON, string(lexer.char), lexer.line)
  case '(':
    return token.New(token.LPAREN, string(lexer.char), lexer.line)
  case ')':
    return token.New(token.RPAREN, string(lexer.char), lexer.line)
  case '{':
    return token.New(token.LBRACE, string(lexer.char), lexer.line)
  case '}':
    return token.New(token.RBRACE, string(lexer.char), lexer.line)
  case '[':
    return token.New(token.LBRACKET, string(lexer.char), lexer.line)
  case ']':
    return token.New(token.RBRACKET, string(lexer.char), lexer.line)
    // Strings
  case '"':
    return lexer.readString()

  default:
    // Identifiers
    if unicode.IsLetter(lexer.char) {
      return lexer.readIdentifier()
    }
    // Numbers
    if unicode.IsNumber(lexer.char) {
      return lexer.readNumber()
    }
    // Illegal
    return token.New(token.ILLEGAL, string(lexer.char), lexer.line)
  }
}

func (lexer *Lexer) peekChar() rune {
  r, _ := utf8.DecodeRuneInString(lexer.source[lexer.position:])

  if r != utf8.RuneError {
    return r
  }

  return 0
}

func (lexer *Lexer) readChar() {
  r, s := utf8.DecodeRuneInString(lexer.source[lexer.position:])

  switch r {
  case utf8.RuneError:
    lexer.char = 0
    return
  default:
    lexer.char = r
    lexer.position += s
  }
}

func (lexer *Lexer) skipWhiteSpace() {
  for checkWhiteSpace(lexer.char) {
    if lexer.char == '\n' {
      lexer.line += 1
    }
    lexer.readChar()
  }
}

func checkWhiteSpace(c rune) bool {
  return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func (lexer *Lexer) readComment() token.Token {
  var comment []rune
  comment = append(comment, lexer.char)

  for {
    lexer.readChar()
    comment = append(comment, lexer.char)

    if lexer.peekChar() == '\n' || lexer.peekChar() == 0 {
      return token.New(token.COMMENT, string(comment), lexer.line)
    }
  }
}

func (lexer *Lexer) readString() token.Token {
  var str []rune
  str = append(str, lexer.char)

  for {
    lexer.readChar()
    str = append(str, lexer.char)

    if lexer.char == '"' {
      return token.New(token.STRING, string(str), lexer.line)
    }
    if lexer.peekChar() == 0 {
      return token.New(token.ILLEGAL, "ILLEFAL", lexer.line)
    }
  }

}

func (lexer *Lexer) readIdentifier() token.Token {
  var ident []rune
  ident = append(ident, lexer.char)

  for unicode.IsLetter(lexer.peekChar()) {
    lexer.readChar()
    ident = append(ident, lexer.char)
  }

  tokenType := token.LookupIdentifier(string(ident))

  return token.New(tokenType, string(ident), lexer.line)
}

func (lexer *Lexer) readNumber() token.Token {
  var number []rune
  number = append(number, lexer.char)

  for unicode.IsNumber(lexer.peekChar()) {
    lexer.readChar()
    number = append(number, lexer.char)
  }
  return token.New(token.NUMBER, string(number), lexer.line)
}
