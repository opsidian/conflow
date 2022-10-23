// Code generated by Conflow. DO NOT EDIT.

package interpreters

import (
	"fmt"

	"github.com/conflowio/conflow/pkg/conflow"
	"github.com/conflowio/conflow/pkg/conflow/annotations"
	"github.com/conflowio/conflow/pkg/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{
				annotations.Type: "configuration",
			},
			ID: "github.com/conflowio/conflow/pkg/schema.Number",
		},
		FieldNames:     map[string]string{"$id": "ID", "const": "Const", "default": "Default", "deprecated": "Deprecated", "description": "Description", "enum": "Enum", "examples": "Examples", "exclusiveMaximum": "ExclusiveMaximum", "exclusiveMinimum": "ExclusiveMinimum", "maximum": "Maximum", "minimum": "Minimum", "multipleOf": "MultipleOf", "nullable": "Nullable", "readOnly": "ReadOnly", "title": "Title", "writeOnly": "WriteOnly", "x-annotations": "Annotations"},
		ParameterNames: map[string]string{"$id": "id", "exclusiveMaximum": "exclusive_maximum", "exclusiveMinimum": "exclusive_minimum", "multipleOf": "multiple_of", "readOnly": "read_only", "writeOnly": "write_only", "x-annotations": "annotations"},
		Properties: map[string]schema.Schema{
			"$id": &schema.String{},
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
				Items: &schema.Any{},
			},
			"exclusiveMaximum": &schema.Number{
				Nullable: true,
			},
			"exclusiveMinimum": &schema.Number{
				Nullable: true,
			},
			"maximum": &schema.Number{
				Nullable: true,
			},
			"minimum": &schema.Number{
				Nullable: true,
			},
			"multipleOf": &schema.Number{
				Nullable: true,
			},
			"nullable":  &schema.Boolean{},
			"readOnly":  &schema.Boolean{},
			"title":     &schema.String{},
			"writeOnly": &schema.Boolean{},
			"x-annotations": &schema.Map{
				AdditionalProperties: &schema.String{},
			},
		},
	})
}

// NumberInterpreter is the Conflow interpreter for the Number block
type NumberInterpreter struct {
}

func (i NumberInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/pkg/schema.Number")
	return s
}

// Create creates a new Number block
func (i NumberInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &schema.Number{}
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i NumberInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i NumberInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *schema.Number
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i NumberInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*schema.Number).ID
	case "const":
		return b.(*schema.Number).Const
	case "default":
		return b.(*schema.Number).Default
	case "deprecated":
		return b.(*schema.Number).Deprecated
	case "description":
		return b.(*schema.Number).Description
	case "enum":
		return b.(*schema.Number).Enum
	case "examples":
		return b.(*schema.Number).Examples
	case "exclusive_maximum":
		return b.(*schema.Number).ExclusiveMaximum
	case "exclusive_minimum":
		return b.(*schema.Number).ExclusiveMinimum
	case "maximum":
		return b.(*schema.Number).Maximum
	case "minimum":
		return b.(*schema.Number).Minimum
	case "multiple_of":
		return b.(*schema.Number).MultipleOf
	case "nullable":
		return b.(*schema.Number).Nullable
	case "read_only":
		return b.(*schema.Number).ReadOnly
	case "title":
		return b.(*schema.Number).Title
	case "write_only":
		return b.(*schema.Number).WriteOnly
	case "annotations":
		return b.(*schema.Number).Annotations
	default:
		panic(fmt.Errorf("unexpected parameter %q in Number", name))
	}
}

func (i NumberInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*schema.Number)
	switch name {
	case "id":
		b.ID = value.(string)
	case "const":
		b.Const = schema.Pointer(value.(float64))
	case "default":
		b.Default = schema.Pointer(value.(float64))
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
		b.ExclusiveMaximum = schema.Pointer(value.(float64))
	case "exclusive_minimum":
		b.ExclusiveMinimum = schema.Pointer(value.(float64))
	case "maximum":
		b.Maximum = schema.Pointer(value.(float64))
	case "minimum":
		b.Minimum = schema.Pointer(value.(float64))
	case "multiple_of":
		b.MultipleOf = schema.Pointer(value.(float64))
	case "nullable":
		b.Nullable = value.(bool)
	case "read_only":
		b.ReadOnly = value.(bool)
	case "title":
		b.Title = value.(string)
	case "write_only":
		b.WriteOnly = value.(bool)
	case "annotations":
		b.Annotations = make(map[string]string, len(value.(map[string]interface{})))
		for valuek, valuev := range value.(map[string]interface{}) {
			b.Annotations[valuek] = valuev.(string)
		}
	}
	return nil
}

func (i NumberInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
