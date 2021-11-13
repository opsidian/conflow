// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
)

// NumberInterpreter is the conflow interpreter for the Number block
type NumberInterpreter struct {
	s schema.Schema
}

func (i NumberInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Metadata: schema.Metadata{
				Annotations: map[string]string{"block.conflow.io/eval_stage": "parse"},
			},
			FieldNames:        map[string]string{"$id": "ID", "annotations": "Annotations", "const": "Const", "default": "Default", "deprecated": "Deprecated", "description": "Description", "enum": "Enum", "examples": "Examples", "exclusiveMaximum": "ExclusiveMaximum", "exclusiveMinimum": "ExclusiveMinimum", "maximum": "Maximum", "minimum": "Minimum", "multipleOf": "MultipleOf", "pointer": "Pointer", "readOnly": "ReadOnly", "title": "Title", "writeOnly": "WriteOnly"},
			JSONPropertyNames: map[string]string{"exclusive_maximum": "exclusiveMaximum", "exclusive_minimum": "exclusiveMinimum", "id": "$id", "multiple_of": "multipleOf", "read_only": "readOnly", "write_only": "writeOnly"},
			Name:              "Number",
			Parameters: map[string]schema.Schema{
				"annotations": &schema.Map{
					AdditionalProperties: &schema.String{},
				},
				"const": &schema.Number{
					Metadata: schema.Metadata{
						Pointer: true,
					},
				},
				"default": &schema.Number{
					Metadata: schema.Metadata{
						Pointer: true,
					},
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
					Metadata: schema.Metadata{
						Pointer: true,
					},
				},
				"exclusive_minimum": &schema.Number{
					Metadata: schema.Metadata{
						Pointer: true,
					},
				},
				"id": &schema.String{},
				"maximum": &schema.Number{
					Metadata: schema.Metadata{
						Pointer: true,
					},
				},
				"minimum": &schema.Number{
					Metadata: schema.Metadata{
						Pointer: true,
					},
				},
				"multiple_of": &schema.Number{
					Metadata: schema.Metadata{
						Pointer: true,
					},
				},
				"pointer":    &schema.Boolean{},
				"read_only":  &schema.Boolean{},
				"title":      &schema.String{},
				"write_only": &schema.Boolean{},
			},
		}
	}
	return i.s
}

// Create creates a new Number block
func (i NumberInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &Number{}
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
	case "pointer":
		return b.(*Number).Pointer
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
	case "pointer":
		b.Pointer = value.(bool)
	case "read_only":
		b.ReadOnly = value.(bool)
	case "title":
		b.Title = value.(string)
	case "write_only":
		b.WriteOnly = value.(bool)
	}
	return nil
}

func (i NumberInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
