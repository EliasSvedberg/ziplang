package evaluator

import (
	"fmt"
	"ziplang/ast"
	"ziplang/object"
	"ziplang/token"
)


var (
  TRUE = &object.Boolean{
    Value: true,
  }
  FALSE = &object.Boolean{
    Value: false,
  }
  NULL = &object.Null{}
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
	case *ast.BooleanExpression:
		return booleanObject(node.Value)
	case *ast.PrefixExpression:
		prefix := &object.Prefix{}
		right := Evaluate(node.Right, environment)

		if isError(right) {
			return right
		}

		prefix.Value = evalPrefixExpression(node.Operator, right)
		return prefix
  case *ast.IdentifierExpression:
    return evalIdentifier(node, environment)
	}

  return newError("TODO: object of type %+v not supported", node)
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

func evalPrefixExpression(operator token.Token, right object.Object) object.Object {
	switch operator.Value {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		// TODO: return error
    return &object.Error{Message: "hehe"}
	}
}

func booleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	} else {
		return FALSE
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.NUMBER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Number).Value
	return &object.Number{Value: -value}
}

func evalIdentifier(node *ast.IdentifierExpression, environment *object.Environment) object.Object {
  if val, ok := environment.Get(node.Value); ok {
    return val
  }

  //if builtin, ok := builtins[node.Value]; ok {
  //  return builtin
  //}

  return newError("identifier not found: " + node.Value)
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}
