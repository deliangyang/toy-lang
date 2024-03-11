package eval

import (
	"github.com/deliangyang/tiny-lang/ast"
	"github.com/deliangyang/tiny-lang/object"
)

// Eval evaluates the AST
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.LetStatement:
		return &object.Integer{Value: node.Value}
	}
	return &object.Null{}
}

func evalStatements(stmts []ast.Statement) object.Object {

	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}

	return result
}

// EvalIntegerExpression evaluates an integer expression
func EvalIntegerExpression(ie *ast.IntegerLiteral) *object.Integer {
	return &object.Integer{Value: ie.Value}
}

// EvalBooleanExpression evaluates a boolean expression
func EvalBooleanExpression(be *ast.Boolean) *object.Boolean {
	return &object.Boolean{Value: be.Value}
}
