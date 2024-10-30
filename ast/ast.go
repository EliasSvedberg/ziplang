package ast

import (
  "ziplang/token"
  "bytes"
  "strconv"
)

type Program struct {
  Statements []Statement
}

func (p *Program) ToString() string {
  var out bytes.Buffer

  out.WriteString("Program {\n")

  for _, s := range p.Statements {
    out.WriteString(s.ToString())
    out.WriteString(",\n")
  }

  out.WriteString("}")

  return out.String()
}

type Node interface {
  TokenValue() string
  ToString() string
}

type Statement interface {
  Node
  StatementNode()
}

type Expression interface {
  Node
  ExpressionNode()
}

type ExpressionStatement struct {
  Token token.Token
  Expression Expression
}

func (es *ExpressionStatement) TokenValue() string {
  return es.Token.Value
}

func (es *ExpressionStatement) ToString() string {
  var out bytes.Buffer

  out.WriteString("ExpressionStatement {\n")
  out.WriteString("Token: ")
  out.WriteString(es.Token.ToString())
  out.WriteString(",\n")
  out.WriteString("Expression: ")
  out.WriteString(es.Expression.ToString())
  out.WriteString(",\n}")
  
  return out.String()
}

func (es *ExpressionStatement) StatementNode() {}

type ReturnStatement struct {
  Token token.Token
  Value Expression
}

func (rs *ReturnStatement) TokenValue() string {
  return rs.Token.Value
}

func (rs *ReturnStatement) ToString() string {
  var out bytes.Buffer

  out.WriteString("ReturnStatement {\n")
  out.WriteString("Token: ")
  out.WriteString(rs.Token.ToString())
  out.WriteString(",\n")
  out.WriteString("Value: ")
  out.WriteString(rs.Value.ToString())
  out.WriteString(",\n}")

  return out.String()
}

func (rs *ReturnStatement) StatementNode() {}

type IdentifierStatement struct {
  Token token.Token
  Type  token.Token // const (::) or var (:=) or reassign var (=)
  Value Expression
}

func (is *IdentifierStatement) TokenValue() string {
  return is.Token.Value
}

func (is *IdentifierStatement) ToString() string {
  var out bytes.Buffer

  out.WriteString("IdentifierStatement {\n")
  out.WriteString("Token: ")
  out.WriteString(is.Token.ToString())
  out.WriteString(",\n")
  out.WriteString("Type: ")
  out.WriteString(is.Type.ToString())
  out.WriteString(",\n")
  out.WriteString("Value: ")
  out.WriteString(is.Value.ToString())
  out.WriteString(",\n}")

  return out.String()
}

func (is *IdentifierStatement) StatementNode() {}


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

type IdentifierExpression struct {
  Token token.Token
  Value string
}

func (ie *IdentifierExpression) TokenValue() string {
  return ie.Token.Value
}

func (ie *IdentifierExpression) ToString() string {
  var out bytes.Buffer

  out.WriteString("IdentifierExpression {\n")
  out.WriteString("Token: ")
  out.WriteString(ie.Token.ToString())
  out.WriteString(",\n")
  out.WriteString("Value: ")
  out.WriteString(ie.Value)
  out.WriteString(",\n")
  out.WriteString("}")

  return out.String()
}

func (ie *IdentifierExpression) ExpressionNode() {}

type StringExpression struct {
  Token token.Token
  Value string
}

func (se *StringExpression) TokenValue() string {
  return se.Token.Value
}

func (se *StringExpression) ToString() string {
  var out bytes.Buffer

  out.WriteString("StringExpression {\n")
  out.WriteString("Token: ")
  out.WriteString(se.Token.ToString())
  out.WriteString(",\n")
  out.WriteString("Value: ")
  out.WriteString(se.Value)
  out.WriteString(",\n")
  out.WriteString("}")

  return out.String()
}

func (se *StringExpression) ExpressionNode() {}

type BooleanExpression struct {
  Token token.Token
  Value string
}

func (be *BooleanExpression) TokenValue() string {
  return be.Token.Value
}

func (be *BooleanExpression) ToString() string {
  var out bytes.Buffer

  out.WriteString("BooleanExpression {\n")
  out.WriteString("Token: ")
  out.WriteString(be.Token.ToString())
  out.WriteString(",\n")
  out.WriteString("Value: ")
  out.WriteString(be.Value)
  out.WriteString(",\n")
  out.WriteString("}")

  return out.String()
}

func (be *BooleanExpression) ExpressionNode() {}

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
