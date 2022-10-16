// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/conflowio/conflow/src/conflow"
	"github.com/conflowio/conflow/src/conflow/annotations"
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{
				annotations.EvalStage: "parse",
				annotations.Type:      "directive",
			},
			ID: "github.com/conflowio/conflow/src/directives.Integer",
		},
		FieldNames:        map[string]string{"$id": "ID", "annotations": "Annotations", "const": "Const", "default": "Default", "deprecated": "Deprecated", "description": "Description", "enum": "Enum", "examples": "Examples", "exclusiveMaximum": "ExclusiveMaximum", "exclusiveMinimum": "ExclusiveMinimum", "format": "Format", "maximum": "Maximum", "minimum": "Minimum", "multipleOf": "MultipleOf", "nullable": "Nullable", "readOnly": "ReadOnly", "title": "Title", "writeOnly": "WriteOnly"},
		JSONPropertyNames: map[string]string{"exclusive_maximum": "exclusiveMaximum", "exclusive_minimum": "exclusiveMinimum", "id": "$id", "multiple_of": "multipleOf", "read_only": "readOnly", "write_only": "writeOnly"},
		ParameterNames:    map[string]string{"$id": "id", "exclusiveMaximum": "exclusive_maximum", "exclusiveMinimum": "exclusive_minimum", "multipleOf": "multiple_of", "readOnly": "read_only", "writeOnly": "write_only"},
		Properties: map[string]schema.Schema{
			"$id": &schema.String{},
			"annotations": &schema.Map{
				AdditionalProperties: &schema.String{},
			},
			"const": &schema.Integer{
				Nullable: true,
			},
			"default": &schema.Integer{
				Nullable: true,
			},
			"deprecated":  &schema.Boolean{},
			"description": &schema.String{},
			"enum": &schema.Array{
				Items: &schema.Integer{},
			},
			"examples": &schema.Array{
				Items: &schema.Any{},
			},
			"exclusiveMaximum": &schema.Integer{
				Nullable: true,
			},
			"exclusiveMinimum": &schema.Integer{
				Nullable: true,
			},
			"format": &schema.String{
				Enum: []string{"int32", "int64"},
			},
			"maximum": &schema.Integer{
				Nullable: true,
			},
			"minimum": &schema.Integer{
				Nullable: true,
			},
			"multipleOf": &schema.Integer{
				Nullable: true,
			},
			"nullable":  &schema.Boolean{},
			"readOnly":  &schema.Boolean{},
			"title":     &schema.String{},
			"writeOnly": &schema.Boolean{},
		},
	})
}

// IntegerInterpreter is the Conflow interpreter for the Integer block
type IntegerInterpreter struct {
}

func (i IntegerInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/directives.Integer")
	return s
}

// Create creates a new Integer block
func (i IntegerInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Integer{}
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i IntegerInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i IntegerInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Integer
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i IntegerInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*Integer).ID
	case "annotations":
		return b.(*Integer).Annotations
	case "const":
		return b.(*Integer).Const
	case "default":
		return b.(*Integer).Default
	case "deprecated":
		return b.(*Integer).Deprecated
	case "description":
		return b.(*Integer).Description
	case "enum":
		return b.(*Integer).Enum
	case "examples":
		return b.(*Integer).Examples
	case "exclusive_maximum":
		return b.(*Integer).ExclusiveMaximum
	case "exclusive_minimum":
		return b.(*Integer).ExclusiveMinimum
	case "format":
		return b.(*Integer).Format
	case "maximum":
		return b.(*Integer).Maximum
	case "minimum":
		return b.(*Integer).Minimum
	case "multiple_of":
		return b.(*Integer).MultipleOf
	case "nullable":
		return b.(*Integer).Nullable
	case "read_only":
		return b.(*Integer).ReadOnly
	case "title":
		return b.(*Integer).Title
	case "write_only":
		return b.(*Integer).WriteOnly
	default:
		panic(fmt.Errorf("unexpected parameter %q in Integer", name))
	}
}

func (i IntegerInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Integer)
	switch name {
	case "id":
		b.ID = value.(string)
	case "annotations":
		b.Annotations = make(map[string]string, len(value.(map[string]interface{})))
		for valuek, valuev := range value.(map[string]interface{}) {
			b.Annotations[valuek] = valuev.(string)
		}
	case "const":
		b.Const = schema.Pointer(value.(int64))
	case "default":
		b.Default = schema.Pointer(value.(int64))
	case "deprecated":
		b.Deprecated = value.(bool)
	case "description":
		b.Description = value.(string)
	case "enum":
		b.Enum = make([]int64, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.Enum[valuek] = valuev.(int64)
		}
	case "examples":
		b.Examples = value.([]interface{})
	case "exclusive_maximum":
		b.ExclusiveMaximum = schema.Pointer(value.(int64))
	case "exclusive_minimum":
		b.ExclusiveMinimum = schema.Pointer(value.(int64))
	case "format":
		b.Format = value.(string)
	case "maximum":
		b.Maximum = schema.Pointer(value.(int64))
	case "minimum":
		b.Minimum = schema.Pointer(value.(int64))
	case "multiple_of":
		b.MultipleOf = schema.Pointer(value.(int64))
	case "nullable":
		b.Nullable = value.(bool)
	case "read_only":
		b.ReadOnly = value.(bool)
	case "title":
		b.Title = value.(string)
	case "write_only":
		b.WriteOnly = value.(bool)
	}
	return nil
}

func (i IntegerInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
