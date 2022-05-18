package object

import (
	"fmt"
	"strings"

	"github.com/josh-weston/go_interpreter/ast"
)

type ObjectType string
type BuiltInFunction func(args ...Object) Object

// note: this works because constants are considered untyped, so
// we can use these untyped constants as ObjecType because there
// will be an implicity conversion when assigned
const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Null struct{} // represents the absence of a value

func (n *Null) Inspect() string  { return NULL_OBJ }
func (n *Null) Type() ObjectType { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERRORL: " + e.Message }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var sb strings.Builder
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	sb.WriteString("fn(")
	sb.WriteString(strings.Join(params, ","))
	sb.WriteString(") {\n")
	sb.WriteString(f.Body.String())
	sb.WriteString("\n")
	return sb.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type Builtin struct {
	Fn BuiltInFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var sb strings.Builder
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	sb.WriteString("[")
	sb.WriteString(strings.Join(elements, ","))
	sb.WriteString("]")
	return sb.String()
}
