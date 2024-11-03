package evaluator

import (
  "testing"
  "ziplang/lexer"
  "ziplang/parser"
  "ziplang/object"
)

func TestEvaluatorNumberExpression(t *testing.T) {
  tests := []struct {
    input string
    expectedOutput interface{}
  }{
    {"5;", 5},
    {"10;", 10},
    {"1337", 1337},
    {"5  ;  ", 5},
  }

  for _, tc := range tests {

    l := lexer.New(tc.input)
    p := parser.New(l)

    program := p.Parse()
    env := object.NewEnvironment()

    evaluatedProgram := Evaluate(program, env)

    result, ok := evaluatedProgram.(*object.Number)

    if !ok {
      t.Errorf("object.Object is not a Number. got=%T (%v)", evaluatedProgram, evaluatedProgram)
    }

    if result.Value != tc.expectedOutput {
      t.Errorf("object has wrong value. got=%d, want=%d", result.Value, tc.expectedOutput)
    }
  }
}
