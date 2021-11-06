// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
)

// ArrayInterpreter is the conflow interpreter for the Array block
type ArrayInterpreter struct {
	s schema.Schema
}

func (i ArrayInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Metadata: schema.Metadata{
				Annotations: map[string]string{"eval_stage": "parse"},
			},
			Name: "Array",
			Properties: map[string]schema.Schema{
				"annotations": &schema.Map{
					AdditionalProperties: &schema.String{},
				},
				"const": &schema.Array{
					Items: &schema.Untyped{},
				},
				"default": &schema.Array{
					Items: &schema.Untyped{},
				},
				"deprecated":  &schema.Boolean{},
				"description": &schema.String{},
				"enum": &schema.Array{
					Items: &schema.Array{
						Items: &schema.Untyped{},
					},
				},
				"examples": &schema.Array{
					Items: &schema.Untyped{},
				},
				"items": &schema.Reference{
					Ref: "http://conflow.schema/github.com/conflowio/conflow/conflow/schema.Schema",
				},
				"max_items": &schema.Integer{
					Metadata: schema.Metadata{
						Pointer: true,
					},
				},
				"min_items":    &schema.Integer{},
				"pointer":      &schema.Boolean{},
				"read_only":    &schema.Boolean{},
				"title":        &schema.String{},
				"unique_items": &schema.Boolean{},
				"write_only":   &schema.Boolean{},
			},
			PropertyNames: map[string]string{"annotations": "Annotations", "const": "Const", "default": "Default", "deprecated": "Deprecated", "description": "Description", "enum": "Enum", "examples": "Examples", "items": "Items", "max_items": "MaxItems", "min_items": "MinItems", "pointer": "Pointer", "read_only": "ReadOnly", "title": "Title", "unique_items": "UniqueItems", "write_only": "WriteOnly"},
			Required:      []string{"items"},
		}
	}
	return i.s
}

// Create creates a new Array block
func (i ArrayInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &Array{}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i ArrayInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i ArrayInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Array
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i ArrayInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "annotations":
		return b.(*Array).Annotations
	case "const":
		return b.(*Array).Const
	case "default":
		return b.(*Array).Default
	case "deprecated":
		return b.(*Array).Deprecated
	case "description":
		return b.(*Array).Description
	case "enum":
		return b.(*Array).Enum
	case "examples":
		return b.(*Array).Examples
	case "max_items":
		return b.(*Array).MaxItems
	case "min_items":
		return b.(*Array).MinItems
	case "pointer":
		return b.(*Array).Pointer
	case "read_only":
		return b.(*Array).ReadOnly
	case "title":
		return b.(*Array).Title
	case "unique_items":
		return b.(*Array).UniqueItems
	case "write_only":
		return b.(*Array).WriteOnly
	default:
		panic(fmt.Errorf("unexpected parameter %q in Array", name))
	}
}

func (i ArrayInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Array)
	switch name {
	case "annotations":
		b.Annotations = make(map[string]string, len(value.(map[string]interface{})))
		for valuek, valuev := range value.(map[string]interface{}) {
			b.Annotations[valuek] = valuev.(string)
		}
	case "const":
		b.Const = value.([]interface{})
	case "default":
		b.Default = value.([]interface{})
	case "deprecated":
		b.Deprecated = value.(bool)
	case "description":
		b.Description = value.(string)
	case "enum":
		b.Enum = make([][]interface{}, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.Enum[valuek] = valuev.([]interface{})
		}
	case "examples":
		b.Examples = value.([]interface{})
	case "max_items":
		b.MaxItems = schema.IntegerPtr(value.(int64))
	case "min_items":
		b.MinItems = value.(int64)
	case "pointer":
		b.Pointer = value.(bool)
	case "read_only":
		b.ReadOnly = value.(bool)
	case "title":
		b.Title = value.(string)
	case "unique_items":
		b.UniqueItems = value.(bool)
	case "write_only":
		b.WriteOnly = value.(bool)
	}
	return nil
}

func (i ArrayInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Array)
	switch name {
	case "items":
		b.Items = value.(schema.Schema)
	}
	return nil
}
