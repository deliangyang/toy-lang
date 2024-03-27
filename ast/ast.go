package ast

import "github.com/deliangyang/tiny-lang/token"

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. '+'
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	return "(" + "todo left" + " " + ie.Operator + " " + "todo right" + ")"
}
