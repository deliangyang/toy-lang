package ast

import "github.com/deliangyang/tiny-lang/token"

type InifxExpression struct {
	Token    token.Token // The operator token, e.g. '+'
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InifxExpression) expressionNode()      {}
func (ie *InifxExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InifxExpression) String() string {
	return "(" + "todo left" + " " + ie.Operator + " " + "todo right" + ")"
}
