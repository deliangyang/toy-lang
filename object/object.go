package object

import "fmt"

// Object is the interface that all objects in the language must implement
type Object interface {
	Type() ObjectType
	Inspect() string
}

// ObjectType is the type of an object
type ObjectType string

const (
	// INTEGER_OBJ is the integer object type
	INTEGER_OBJ = "INTEGER"
	// BOOLEAN_OBJ is the boolean object type
	BOOLEAN_OBJ = "BOOLEAN"
	// NULL_OBJ is the null object type
	NULL_OBJ = "NULL"
)

// Integer is the integer object
type Integer struct {
	Value int64
}

// Inspect returns the string representation of the integer object
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type returns the type of the integer object
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Boolean is the boolean object
type Boolean struct {
	Value bool
}

// Inspect returns the string representation of the boolean object
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type returns the type of the boolean object
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Null is the null object
type Null struct{}

// Inspect returns the string representation of the null object
func (n *Null) Inspect() string {
	return "null"
}

// Type returns the type of the null object
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// ReturnValue is the return value object
type ReturnValue struct {
	Value Object
}

// Inspect returns the string representation of the return value object
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Type returns the type of the return value object
func (rv *ReturnValue) Type() ObjectType {
	return ""
}

// Error is the error object
type Error struct {
	Message string
}

// Inspect returns the string representation of the error object
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Type returns the type of the error object
func (e *Error) Type() ObjectType {
	return ""
}
