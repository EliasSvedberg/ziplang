package object

import (
	"strconv"
)

const (
	NUMBER_OBJ       = "NUMBER"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	BOOLEAN_OBJ      = "BOOLEAN"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	ToString() string
}

type Number struct {
	Value int
}

func (n *Number) Type() ObjectType {
	return NUMBER_OBJ
}

func (n *Number) ToString() string {
	return strconv.Itoa(n.Value)
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return NUMBER_OBJ
}

func (s *String) ToString() string {
	return s.Value
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) ToString() string {
	return rv.Value.ToString()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) ToString() string {
	return "ERROR: " + e.Message
}

type Boolean struct {
  Value bool
}

func (b *Boolean) Type() ObjectType {
  return BOOLEAN_OBJ
}

func (b *Boolean) ToString() string {
  if (b.Value) { return "true" } else { return "false" }
}
