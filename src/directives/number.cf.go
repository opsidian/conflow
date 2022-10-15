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
			ID: "github.com/conflowio/conflow/src/directives.Number",
		},
		FieldNames:        map[string]string{"$id": "ID", "annotations": "Annotations", "const": "Const", "default": "Default", "deprecated": "Deprecated", "description": "Description", "enum": "Enum", "examples": "Examples", "exclusiveMaximum": "ExclusiveMaximum", "exclusiveMinimum": "ExclusiveMinimum", "maximum": "Maximum", "minimum": "Minimum", "multipleOf": "MultipleOf", "nullable": "Nullable", "readOnly": "ReadOnly", "title": "Title", "writeOnly": "WriteOnly"},
		JSONPropertyNames: map[string]string{"exclusive_maximum": "exclusiveMaximum", "exclusive_minimum": "exclusiveMinimum", "id": "$id", "multiple_of": "multipleOf", "read_only": "readOnly", "write_only": "writeOnly"},
		Name:              "Number",
		Parameters: map[string]schema.Schema{
			"annotations": &schema.Map{
				AdditionalProperties: &schema.String{},
			},
			"const": &schema.Number{
				Nullable: true,
			},
			"default": &schema.Number{
				Nullable: true,
			},
			"deprecated":  &schema.Boolean{},
			"description": &schema.String{},
			"enum": &schema.Array{
				Items: &schema.Number{},
			},
			"examples": &schema.Array{
				Items: &schema.Untyped{},
			},
			"exclusive_maximum": &schema.Number{
				Nullable: true,
			},
			"exclusive_minimum": &schema.Number{
				Nullable: true,
			},
			"id": &schema.String{},
			"maximum": &schema.Number{
				Nullable: true,
			},
			"minimum": &schema.Number{
				Nullable: true,
			},
			"multiple_of": &schema.Number{
				Nullable: true,
			},
			"nullable":   &schema.Boolean{},
			"read_only":  &schema.Boolean{},
			"title":      &schema.String{},
			"write_only": &schema.Boolean{},
		},
	})
}

// NumberInterpreter is the Conflow interpreter for the Number block
type NumberInterpreter struct {
}

func (i NumberInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/directives.Number")
	return s
}

// Create creates a new Number block
func (i NumberInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Number{}
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i NumberInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i NumberInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Number
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i NumberInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "annotations":
		return b.(*Number).Annotations
	case "const":
		return b.(*Number).Const
	case "default":
		return b.(*Number).Default
	case "deprecated":
		return b.(*Number).Deprecated
	case "description":
		return b.(*Number).Description
	case "enum":
		return b.(*Number).Enum
	case "examples":
		return b.(*Number).Examples
	case "exclusive_maximum":
		return b.(*Number).ExclusiveMaximum
	case "exclusive_minimum":
		return b.(*Number).ExclusiveMinimum
	case "id":
		return b.(*Number).ID
	case "maximum":
		return b.(*Number).Maximum
	case "minimum":
		return b.(*Number).Minimum
	case "multiple_of":
		return b.(*Number).MultipleOf
	case "nullable":
		return b.(*Number).Nullable
	case "read_only":
		return b.(*Number).ReadOnly
	case "title":
		return b.(*Number).Title
	case "write_only":
		return b.(*Number).WriteOnly
	default:
		panic(fmt.Errorf("unexpected parameter %q in Number", name))
	}
}

func (i NumberInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Number)
	switch name {
	case "annotations":
		b.Annotations = make(map[string]string, len(value.(map[string]interface{})))
		for valuek, valuev := range value.(map[string]interface{}) {
			b.Annotations[valuek] = valuev.(string)
		}
	case "const":
		b.Const = schema.NumberPtr(value.(float64))
	case "default":
		b.Default = schema.NumberPtr(value.(float64))
	case "deprecated":
		b.Deprecated = value.(bool)
	case "description":
		b.Description = value.(string)
	case "enum":
		b.Enum = make([]float64, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.Enum[valuek] = valuev.(float64)
		}
	case "examples":
		b.Examples = value.([]interface{})
	case "exclusive_maximum":
		b.ExclusiveMaximum = schema.NumberPtr(value.(float64))
	case "exclusive_minimum":
		b.ExclusiveMinimum = schema.NumberPtr(value.(float64))
	case "id":
		b.ID = value.(string)
	case "maximum":
		b.Maximum = schema.NumberPtr(value.(float64))
	case "minimum":
		b.Minimum = schema.NumberPtr(value.(float64))
	case "multiple_of":
		b.MultipleOf = schema.NumberPtr(value.(float64))
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

func (i NumberInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
