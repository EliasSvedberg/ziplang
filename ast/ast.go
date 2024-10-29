package ast

import (
  "ziplang/token"
  "bytes"
  "strconv"
)

type Program struct {
  Expressions []Expression
}

func (p *Program) ToString() string {
  var out bytes.Buffer

  out.WriteString("Program {\n")

  for _, e := range p.Expressions {
    out.WriteString(e.ToString())
    out.WriteString(",\n")
  }

  out.WriteString("}")

  return out.String()
}

type Expression interface {
  TokenValue() string
  ToString() string
  ExpressionNode()
}

type NumberExpression struct {
  Token token.Token
  Value int
}

func (ne *NumberExpression) TokenValue() string {
  return ne.Token.Value
}

func (ne *NumberExpression) ToString() string {
  var out bytes.Buffer

  out.WriteString("NumberExpression {\n")
  out.WriteString("Token: ")
  out.WriteString(ne.Token.ToString())
  out.WriteString(",\n")
  out.WriteString("Value: ")
  out.WriteString(strconv.Itoa(ne.Value))
  out.WriteString(",\n")
  out.WriteString("}")

  return out.String()
}

func (ne *NumberExpression) ExpressionNode() {}

type InfixExpression struct {
  Token token.Token
  Left Expression
  Operator token.Token
  Right Expression
}

func (ie *InfixExpression) TokenValue() string {
  return ie.Token.Value
}

func (ie *InfixExpression) ToString() string {
  var out bytes.Buffer

  out.WriteString("InfixExpression {\n")
  out.WriteString("Token: ")
  out.WriteString(ie.Token.ToString())
  out.WriteString(",\n")
  out.WriteString("Left: ")
  out.WriteString(ie.Left.ToString())
  out.WriteString(",\n")
  out.WriteString("Operator: ")
  out.WriteString(ie.Operator.ToString())
  out.WriteString(",\n")
  out.WriteString("Right: ")
  out.WriteString(ie.Right.ToString())
  out.WriteString(",\n")
  out.WriteString("}")

  return out.String()
}

func (ie *InfixExpression) ExpressionNode() {}
