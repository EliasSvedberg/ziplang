package object

import (
  "bytes"
  "strconv"
)

const (
  NUMBER_OBJ = "NUMBER"
  RETURN_VALUE_OBJ = "RETURN_VALUE"
  ERROR_OBJ = "ERROR"
)

type ObjectType string

type Object interface {
  Type() ObjectType
  ToString() string
}

type Number struct {
  Value int
}

func (i *Number) Type() ObjectType {
  return NUMBER_OBJ
}

func (i *Number) ToString() string {
  var out bytes.Buffer
  out.WriteString(strconv.Itoa(i.Value))
  return out.String()
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
