package ast

import (
	"reflect"
	"strings"
	"testing"
	"ziplang/token"
)

func TestAstNodeTokenValue(t *testing.T) {
	tests := []struct {
		input          Node
		expectedOutput string
	}{
		{&NumberExpression{
			Token: token.Token{
				Type:  token.NUMBER,
				Value: "1",
				Line:  1,
			},
			Value: 1,
		}, "1"},
		{&StringExpression{
			Token: token.Token{
				Type:  token.STRING,
				Value: "hello",
				Line:  1,
			},
			Value: "hello",
		}, "hello"},
		{&IdentifierExpression{
			Token: token.Token{
				Type:  token.IDENTIFIER,
				Value: "foo",
				Line:  1,
			},
			Value: "foo",
		}, "foo"},
	}

	for _, tc := range tests {
		result := tc.input.TokenValue()

		if result != tc.expectedOutput {
			t.Errorf("TokenValue() failed for the node type: %+v.\nExpected: %s\nGot:\n%s", reflect.TypeOf(result), tc.expectedOutput, result)
		}
	}
}

func TestAstNodeToString(t *testing.T) {
	tests := []struct {
		input          Node
		expectedOutput string
	}{
		{&NumberExpression{
			Token: token.Token{
				Type:  token.NUMBER,
				Value: "1",
				Line:  1,
			},
			Value: 1,
		}, `NumberExpression {
      Token: Token {
        Type: NUMBER,
        Value: 1,
        Line: 1,
      },
      Value: 1,
    }`},
		{&StringExpression{
			Token: token.Token{
				Type:  token.STRING,
				Value: "hello",
				Line:  1,
			},
			Value: "hello",
		}, `StringExpression {
      Token: Token {
        Type: STRING,
        Value: hello,
        Line: 1,
      },
      Value: hello,
    }`},
		{&IdentifierExpression{
			Token: token.Token{
				Type:  token.IDENTIFIER,
				Value: "foo",
				Line:  1,
			},
			Value: "foo",
		}, `IdentifierExpression {
      Token: Token {
        Type: IDENTIFIER,
        Value: foo,
        Line: 1,
      },
      Value: foo,
    }`},
	}

	for _, tc := range tests {
		result := tc.input.ToString()

		if strings.ReplaceAll(result, " ", "") != strings.ReplaceAll(tc.expectedOutput, " ", "") {
			t.Errorf("ToString() failed for the node type: %+v.\nExpected: %s\nGot:\n%s", reflect.TypeOf(result), strings.ReplaceAll(tc.expectedOutput, " ", ""), strings.ReplaceAll(result, " ", ""))
		}
	}
}
