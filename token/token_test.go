package token

import (
  "testing"
  "strings"
)

func TestTokenLookupIdentifier(t *testing.T) {
  tests := []struct {
    input string
    expectedType TokenType
  }{
    {"foo", IDENTIFIER},
    {"bar", IDENTIFIER},
    {"baz", IDENTIFIER},
    {"return", RETURN},
    {"false", FALSE},
    {"true", TRUE},
  }

  for _, tc := range tests {
    if (LookupIdentifier(tc.input) != tc.expectedType) {
      t.Errorf("Wrong TokenType. Expected: %T\nGot: %T", LookupIdentifier(tc.input), tc.expectedType)
    }
  }
}

func TestTokenToString(t *testing.T) {
  tests := []struct {
    input Token
    expectedOutput string
  }{
    {Token {
      Type: IDENTIFIER,
      Value: "foo",
      Line: 1,
    }, `Token {
      Type: IDENTIFIER,
      Value: foo,
      Line: 1,
    }`},
    {Token {
      Type: STRING,
      Value: "\"foo\"",
      Line: 1,
    }, `Token {
      Type: STRING,
      Value: "foo",
      Line: 1,
    }`},
    {Token {
      Type: NUMBER,
      Value: "3",
      Line: 1,
    }, `Token {
      Type: NUMBER,
      Value: 3,
      Line: 1,
    }`},
    {Token {
      Type: PLUS,
      Value: "+",
      Line: 1,
    }, `Token {
      Type: PLUS,
      Value: +,
      Line: 1,
    }`},
    {Token {
      Type: MINUS,
      Value: "-",
      Line: 1,
    }, `Token {
      Type: MINUS,
      Value: -,
      Line: 1,
    }`},
    {Token {
      Type: EOF,
      Value: "EOF",
      Line: 1,
    }, `Token {
      Type: EOF,
      Value: EOF,
      Line: 1,
    }`},
  }

  for _, tc := range tests {

		if strings.ReplaceAll(tc.input.ToString(), " ", "") != strings.ReplaceAll(tc.expectedOutput, " ", "") {
			t.Errorf("ToString() failed for the node type: %+v.\nExpected: %s\nGot:\n%s", tc.input.Type, strings.ReplaceAll(tc.expectedOutput, " ", ""), strings.ReplaceAll(tc.input.ToString(), " ", ""))
		}
  }
}

