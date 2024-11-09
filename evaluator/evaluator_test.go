package evaluator

import (
	"testing"
	"ziplang/lexer"
	"ziplang/object"
	"ziplang/parser"
)

func TestEvaluatorPrefixExpression(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput interface{}
	}{
		{"-1;", -1},
		{"-true;", false},
		{"!false", true},
		{"!true", false},
	}

	for _, tc := range tests {
		l := lexer.New(tc.input)
		p := parser.New(l)

		program := p.Parse()
		env := object.NewEnvironment()

		evaluatedProgram := Evaluate(program, env)

		wrappingResult, ok := evaluatedProgram.(*object.Prefix)

		if !ok {
			t.Errorf("object.Object is not a Prefix. got=%T (%v)", evaluatedProgram, evaluatedProgram)
		}

		switch wrappingResult.Value.Type() {
		case object.NUMBER_OBJ:

			result, ok := wrappingResult.Value.(*object.Number)

			if !ok {
				t.Errorf("object.Object is not a Number. got=%T (%v)", evaluatedProgram, evaluatedProgram)
			}

			if result.Value != tc.expectedOutput {
				t.Errorf("object has wrong value. got=%+v, want=%+v", result.Value, tc.expectedOutput)
			}

		case object.BOOLEAN_OBJ:
			result, ok := wrappingResult.Value.(*object.Boolean)

			if !ok {
				t.Errorf("object.Object is not a Boolean. got=%T (%v)", evaluatedProgram, evaluatedProgram)
			}

			if result.Value != tc.expectedOutput {
				t.Errorf("object has wrong value. got=%+v, want=%+v", result.Value, tc.expectedOutput)
			}

		case object.ERROR_OBJ:
			_, ok := wrappingResult.Value.(*object.Error)

			if !ok {
				t.Errorf("object.Object is not a Error. got=%T (%v)", evaluatedProgram, evaluatedProgram)
			}

		default:
			t.Errorf("Unsupporeted value for prefixexpression. got=%+v, want=Boolean || Number", wrappingResult.Value.Type())
		}
	}
}

func TestEvaluatorBooleanExpression(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput interface{}
	}{
		{"true;", true},
		{"false;", false},
		{"false  ", false},
		{"true  ", true},
	}

	for _, tc := range tests {
		l := lexer.New(tc.input)
		p := parser.New(l)

		program := p.Parse()
		env := object.NewEnvironment()

		evaluatedProgram := Evaluate(program, env)

		result, ok := evaluatedProgram.(*object.Boolean)

		if !ok {
			t.Errorf("object.Object is not a Boolean. got=%T (%v)", evaluatedProgram, evaluatedProgram)
		}

		if result.Value != tc.expectedOutput {
			t.Errorf("object has wrong value. got=%+v, want=%+v", result.Value, tc.expectedOutput)
		}
	}
}

func TestEvalatorStringExpression(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput interface{}
	}{
		{`"teststring"`, `"teststring"`},
		{`"asd"`, `"asd"`},
		{`"333"`, `"333"`},
	}

	for _, tc := range tests {
		l := lexer.New(tc.input)
		p := parser.New(l)

		program := p.Parse()
		env := object.NewEnvironment()

		evaluatedProgram := Evaluate(program, env)

		result, ok := evaluatedProgram.(*object.String)

		if !ok {
			t.Errorf("object.Object is not a String. got=%T (%v)", evaluatedProgram, evaluatedProgram)
		}

		if result.Value != tc.expectedOutput {
			t.Errorf("object has wrong value. got=%s, want=%s", result.Value, tc.expectedOutput)
		}
	}
}

func TestEvaluatorNumberExpression(t *testing.T) {
	tests := []struct {
		input          string
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
