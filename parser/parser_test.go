package parser

import (
	"strings"
	"testing"
	"ziplang/lexer"
)

func TestParserInfixExpresion(t *testing.T) {
  tests := []struct {
    input string
    expectedProgram string
  }{
		{"1 != 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
          Token: Token {
            Type: NOT_EQ,
            Value: !=,
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
            Type: NOT_EQ,
            Value: !=,
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
      },
    }`},
		{"1 == 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
          Token: Token {
            Type: EQ,
            Value: ==,
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
            Type: EQ,
            Value: ==,
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
      },
    }`},
		{"1 > 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
          Token: Token {
            Type: GT,
            Value: >,
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
            Type: GT,
            Value: >,
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
      },
    }`},
		{"1 < 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
          Token: Token {
            Type: LT,
            Value: <,
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
            Type: LT,
            Value: <,
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
      },
    }`},
		{"1 / 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
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

func TestParserFunctionExpression(t *testing.T) {
  tests := []struct {
    input string
    expectedProgram string
  }{
    {"fn(a) { return a; }",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: FUNCTION,
          Value: fn,
          Line: 1,
        },
        Expression: FunctionExpression {
          Token: Token{
            Type: FUNCTION,
            Value: fn,
            Line: 1,
          },
          Parameters: IdentifierExpression {
            Token: Token {
              Type: IDENTIFIER,
              Value: a,
              Line: 1,
            },
            Value: a,
          },
          Body: BlockStatement {
            Token: Token {
              Type: LBRACE,
              Value: {,
              Line: 1,
            },
            Statements: ReturnStatement {
              Token: Token {
                Type: RETURN,
                Value: return,
                Line: 1,
              },
              Value: IdentifierExpression {
                Token: Token {
                  Type: IDENTIFIER,
                  Value: a,
                  Line: 1,
                },
                Value: a,
              },
            },
          },
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

func TestParserGroupedExpression(t *testing.T) {
  tests := []struct {
    input string
    expectedProgram string
  }{
    {"(1 + 1) * 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: LPAREN,
          Value: (,
          Line: 1,
        },
        Expression: InfixExpression {
          Token: Token {
            Type: ASTERISK,
            Value: *,
            Line: 1,
          },
          Left: InfixExpression {
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
      },
    }`},
    {"(1 - 1) * 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: LPAREN,
          Value: (,
          Line: 1,
        },
        Expression: InfixExpression {
          Token: Token {
            Type: ASTERISK,
            Value: *,
            Line: 1,
          },
          Left: InfixExpression {
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

func TestParserPrefixExpression(t *testing.T) {
  tests := []struct {
    input string
    expectedProgram string
  }{
    {"-1;",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: MINUS,
          Value: -,
          Line: 1,
        },
        Expression: PrefixExpression {
          Token: Token {
            Type: MINUS,
            Value: -,
            Line: 1,
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
      },
    }`},
    {"!1;",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: BANG,
          Value: !,
          Line: 1,
        },
        Expression: PrefixExpression {
          Token: Token {
            Type: BANG,
            Value: !,
            Line: 1,
          },
          Operator: Token {
            Type: BANG,
            Value: !,
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

func TestParserIdentifierStatement(t *testing.T) {
  tests := []struct {
    input string
    expectedProgram string
  }{
    {"foo ::  1;",
    `Program {
      IdentifierStatement {
        Token: Token {
          Type: IDENTIFIER,
          Value: foo,
          Line: 1,
        },
        Type: Token {
          Type: CONST,
          Value: ::,
          Line: 1,
        },
        Value: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
      },
    }`},
    {"bar :=  1;",
    `Program {
      IdentifierStatement {
        Token: Token {
          Type: IDENTIFIER,
          Value: bar,
          Line: 1,
        },
        Type: Token {
          Type: VAR,
          Value: :=,
          Line: 1,
        },
        Value: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
      },
    }`},
    {"baz =  1;",
    `Program {
      IdentifierStatement {
        Token: Token {
          Type: IDENTIFIER,
          Value: baz,
          Line: 1,
        },
        Type: Token {
          Type: ASSIGN,
          Value: =,
          Line: 1,
        },
        Value: NumberExpression {
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

func TestParserReturnStatement(t *testing.T) {
  tests := []struct {
    input string
    expectedProgram string
  }{
		{"return  1;",
    `Program {
      ReturnStatement {
        Token: Token {
          Type: RETURN,
          Value: return,
          Line: 1,
        },
        Value: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
      },
    }`},
		{"return  true;",
    `Program {
      ReturnStatement {
        Token: Token {
          Type: RETURN,
          Value: return,
          Line: 1,
        },
        Value: BooleanExpression {
          Token: Token {
            Type: TRUE,
            Value: true,
            Line: 1,
          },
          Value: true,
        },
      },
    }`},
		{"return  \"foo\";",
    `Program {
      ReturnStatement {
        Token: Token {
          Type: RETURN,
          Value: return,
          Line: 1,
        },
        Value: StringExpression {
          Token: Token {
            Type: STRING,
            Value: "foo",
            Line: 1,
          },
          Value: "foo",
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

func TestParserBooleanExpression(t *testing.T) {
  tests := []struct {
    input string
    expectedProgram string
  }{
		{"true  ;",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: TRUE,
          Value: true,
          Line: 1,
        },
        Expression: BooleanExpression {
          Token: Token {
            Type: TRUE,
            Value: true,
            Line: 1,
          },
          Value: true,
        },
      },
    }`},
		{"false;",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: FALSE,
          Value: false,
          Line: 1,
        },
        Expression: BooleanExpression {
          Token: Token {
            Type: FALSE,
            Value: false,
            Line: 1,
          },
          Value: false,
        },
      },
    }`},
		{"false",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: FALSE,
          Value: false,
          Line: 1,
        },
        Expression: BooleanExpression {
          Token: Token {
            Type: FALSE,
            Value: false,
            Line: 1,
          },
          Value: false,
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

func TestParserStringExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedProgram string
	}{
    {"\"foo\";",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: STRING,
          Value: "foo",
          Line: 1,
        },
        Expression: StringExpression {
          Token: Token {
            Type: STRING,
            Value: "foo",
            Line: 1,
          },
          Value: "foo",
        },
      },
    }`},
    {"\"bar\";",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: STRING,
          Value: "bar",
          Line: 1,
        },
        Expression: StringExpression {
          Token: Token {
            Type: STRING,
            Value: "bar",
            Line: 1,
          },
          Value: "bar",
        },
      },
    }`},
    {"\"baz\";",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: STRING,
          Value: "baz",
          Line: 1,
        },
        Expression: StringExpression {
          Token: Token {
            Type: STRING,
            Value: "baz",
            Line: 1,
          },
          Value: "baz",
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

func TestParserIdentifierExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedProgram string
	}{
    {"foo;",
    `Program {
      ExpressionStatement { 
        Token: Token {
          Type: IDENTIFIER,
          Value: foo,
          Line: 1,
        },
        Expression: IdentifierExpression {
          Token: Token {
            Type: IDENTIFIER,
            Value: foo,
            Line: 1,
          },
          Value: foo,
        },
      },
    }`},
    {"  bar;",
    `Program {
      ExpressionStatement { 
        Token: Token {
          Type: IDENTIFIER,
          Value: bar,
          Line: 1,
        },
        Expression: IdentifierExpression {
          Token: Token {
            Type: IDENTIFIER,
            Value: bar,
            Line: 1,
          },
          Value: bar,
        },
      },
    }`},
    {"  baz  ;",
    `Program {
      ExpressionStatement { 
        Token: Token {
          Type: IDENTIFIER,
          Value: baz,
          Line: 1,
        },
        Expression: IdentifierExpression {
          Token: Token {
            Type: IDENTIFIER,
            Value: baz,
            Line: 1,
          },
          Value: baz,
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

func TestParserMultiExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedProgram string
	}{
    {"4-1; 2 * 2;",
    `Program {
      ExpressionStatement {
        Token: Token {
          Type: NUMBER,
          Value: 4,
          Line: 1,
        },
        Expression: InfixExpression {
          Token: Token {
            Type: MINUS,
            Value: -,
            Line: 1,
          },
          Left: NumberExpression {
            Token: Token {
              Type: NUMBER,
              Value: 4,
              Line: 1,
            },
            Value: 4,
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
      },
      ExpressionStatement {
        Token: Token {
          Type: NUMBER,
          Value: 2,
          Line: 1,
        },
        Expression: InfixExpression {
          Token: Token {
            Type: ASTERISK,
            Value: *,
            Line: 1,
          },
          Left: NumberExpression {
            Token: Token {
              Type: NUMBER,
              Value: 2,
              Line: 1,
            },
            Value: 2,
          },
          Operator: Token {
            Type: ASTERISK,
            Value: *,
            Line: 1,
          },
          Right: NumberExpression {
            Token: Token {
              Type: NUMBER,
              Value: 2,
              Line: 1,
            },
            Value: 2,
          },
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

func TestParserLineNumbers(t *testing.T) {
	tests := []struct {
		input           string
		expectedProgram string
	}{
		{`
    1;
    `,
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 2,
        },
        Expression: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 2,
          },
          Value: 1,
        },
      },
    }`},
		{`


    1337;
    `,
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1337,
          Line: 4,
        },
        Expression: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1337,
            Line: 4,
          },
          Value: 1337,
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

func TestParserGeneric(t *testing.T) {
	tests := []struct {
		input           string
		expectedProgram string
	}{
    {"1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
      },
    }`},
    {"(1);",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: LPAREN,
          Value: (,
          Line: 1,
        },
        Expression: NumberExpression {
          Token: Token {
            Type: NUMBER,
            Value: 1,
            Line: 1,
          },
          Value: 1,
        },
      },
    }`},
		{"1 % 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
          Token: Token {
            Type: MODULO,
            Value: %,
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
            Type: MODULO,
            Value: %,
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
      },
    }`},
		{"1 + 1",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
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
      },
    }`},
		{"1 - 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
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
      },
    }`},
    {"1 * 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
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
      },
    }`},
		{"1 / 1;",
    `Program {
      ExpressionStatement {
        Token: Token{
          Type: NUMBER,
          Value: 1,
          Line: 1,
        },
        Expression: InfixExpression {
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
