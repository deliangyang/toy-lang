package eval

import (
	"github.com/deliangyang/tiny-lang/ast"
	"github.com/deliangyang/tiny-lang/object"
)

// Eval evaluates the AST
func Eval(node ast.Node, env object.Enviroment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		env.Set(node.Name.Value, val)
		return val
	case *ast.Identifier:
		return evalIdentifier(node, env)
	}
	return &object.Null{}
}

func evalIdentifier(node *ast.Identifier, env object.Enviroment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return &object.Null{}
}

func evalStatements(stmts []ast.Statement, env object.Enviroment) object.Object {

	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt, env)
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
