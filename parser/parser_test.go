package parser

import (
	"strings"
	"testing"
	"ziplang/lexer"
)

func TestParserLineNumbers(t *testing.T) {
	tests := []struct {
		input           string
		expectedProgram string
	}{
		{`
    1;
    `,
			`Program {
      NumberExpression {
        Token: Token {
          Type: NUMBER,
          Value: 1,
          Line: 2,
        },
        Value: 1,
      },
    }`},
		{`


    1337;
    `,
			`Program {
      NumberExpression {
        Token: Token {
          Type: NUMBER,
          Value: 1337,
          Line: 4,
        },
        Value: 1337,
      },
    }`},
	}

	for _, tc := range tests {
		l := lexer.New(tc.input)

		p := New(l)

		program := p.Parse()

		msg, hasErrors := p.ReportParserErrors()
		if hasErrors != nil {
			t.Errorf(msg)
		}

		if strings.ReplaceAll(program.ToString(), " ", "") != strings.ReplaceAll(tc.expectedProgram, " ", "") {
			t.Errorf("wrong program generated. Expected:\n%s\ngot:\n%s", strings.ReplaceAll(tc.expectedProgram, " ", ""), strings.ReplaceAll(program.ToString(), " ", ""))
		}
	}
}

func TestParserGeneric(t *testing.T) {
	tests := []struct {
		input           string
		expectedProgram string
	}{
		{"1;",
			`Program {
      NumberExpression {
        Token: Token {
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Value: 1,
      },
    }`},
		{"1 + 1",
			`Program {
      InfixExpression {
        Token: Token {
          Type: PLUS,
          Value: +,
          Line: 1,
        },
        Left: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
        Operator: Token {
          Type: PLUS,
          Value: +,
          Line: 1,
        },
        Right: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
      },
    }`},
		{"1 - 1;",
			`Program {
      InfixExpression {
        Token: Token {
          Type: MINUS,
          Value: -,
          Line: 1,
        },
        Left: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
        Operator: Token {
          Type: MINUS,
          Value: -,
          Line: 1,
        },
        Right: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
      },
    }`},
		{"1 * 1;",
			`Program {
      InfixExpression {
        Token: Token {
          Type: ASTERISK,
          Value: *,
          Line: 1,
        },
        Left: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
        Operator: Token {
          Type: ASTERISK,
          Value: *,
          Line: 1,
        },
        Right: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
      },
    }`},
		{"1 / 1;",
			`Program {
      InfixExpression {
        Token: Token {
          Type: SLASH,
          Value: /,
          Line: 1,
        },
        Left: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
        Operator: Token {
          Type: SLASH,
          Value: /,
          Line: 1,
        },
        Right: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
      },
    }`},
	}

	for _, tc := range tests {
		l := lexer.New(tc.input)

		p := New(l)

		program := p.Parse()

		msg, hasErrors := p.ReportParserErrors()
		if hasErrors != nil {
			t.Errorf(msg)
		}

		if strings.ReplaceAll(program.ToString(), " ", "") != strings.ReplaceAll(tc.expectedProgram, " ", "") {
			t.Errorf("wrong program generated. Expected:\n%s\ngot:\n%s", strings.ReplaceAll(tc.expectedProgram, " ", ""), strings.ReplaceAll(program.ToString(), " ", ""))
		}
	}
}
