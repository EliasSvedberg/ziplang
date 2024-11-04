package evaluator

import (
	"ziplang/ast"
	"ziplang/object"
)

func Evaluate(node ast.Node, environment *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evaluateProgram(node.Statements, environment)
	case *ast.ExpressionStatement:
		return Evaluate(node.Expression, environment)
	case *ast.NumberExpression:
		return &object.Number{
			Value: node.Value,
		}
	case *ast.StringExpression:
		return &object.String{
			Value: node.Value,
		}
	}

	return nil
}

func evaluateProgram(statements []ast.Statement, environment *object.Environment) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Evaluate(statement, environment)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}
