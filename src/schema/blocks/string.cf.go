// Code generated by Conflow. DO NOT EDIT.

package blocks

import (
	"fmt"
	"github.com/conflowio/conflow/src/conflow"
	"github.com/conflowio/conflow/src/schema"
	"github.com/conflowio/conflow/src/schema/formats"
	"regexp"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{"block.conflow.io/type": "configuration"},
			ID:          "github.com/conflowio/conflow/src/schema/blocks.String",
		},
		FieldNames:        map[string]string{"$id": "ID", "annotations": "Annotations", "const": "Const", "default": "Default", "deprecated": "Deprecated", "description": "Description", "enum": "Enum", "examples": "Examples", "format": "Format", "maxLength": "MaxLength", "minLength": "MinLength", "nullable": "Nullable", "pattern": "Pattern", "readOnly": "ReadOnly", "title": "Title", "writeOnly": "WriteOnly"},
		JSONPropertyNames: map[string]string{"id": "$id", "max_length": "maxLength", "min_length": "minLength", "read_only": "readOnly", "write_only": "writeOnly"},
		Name:              "String",
		Parameters: map[string]schema.Schema{
			"annotations": &schema.Map{
				AdditionalProperties: &schema.String{},
			},
			"const": &schema.String{
				Nullable: true,
			},
			"default": &schema.String{
				Nullable: true,
			},
			"deprecated":  &schema.Boolean{},
			"description": &schema.String{},
			"enum": &schema.Array{
				Items: &schema.String{},
			},
			"examples": &schema.Array{
				Items: &schema.Untyped{},
			},
			"format": &schema.String{},
			"id":     &schema.String{},
			"max_length": &schema.Integer{
				Nullable: true,
			},
			"min_length": &schema.Integer{},
			"nullable":   &schema.Boolean{},
			"pattern": &schema.String{
				Format:   "regex",
				Nullable: true,
			},
			"read_only":  &schema.Boolean{},
			"title":      &schema.String{},
			"write_only": &schema.Boolean{},
		},
	})
}

// StringInterpreter is the Conflow interpreter for the String block
type StringInterpreter struct {
}

func (i StringInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/schema/blocks.String")
	return s
}

// Create creates a new String block
func (i StringInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &String{}
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i StringInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i StringInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *String
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i StringInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "annotations":
		return b.(*String).Annotations
	case "const":
		return b.(*String).Const
	case "default":
		return b.(*String).Default
	case "deprecated":
		return b.(*String).Deprecated
	case "description":
		return b.(*String).Description
	case "enum":
		return b.(*String).Enum
	case "examples":
		return b.(*String).Examples
	case "format":
		return b.(*String).Format
	case "id":
		return b.(*String).ID
	case "max_length":
		return b.(*String).MaxLength
	case "min_length":
		return b.(*String).MinLength
	case "nullable":
		return b.(*String).Nullable
	case "pattern":
		return b.(*String).Pattern
	case "read_only":
		return b.(*String).ReadOnly
	case "title":
		return b.(*String).Title
	case "write_only":
		return b.(*String).WriteOnly
	default:
		panic(fmt.Errorf("unexpected parameter %q in String", name))
	}
}

func (i StringInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*String)
	switch name {
	case "annotations":
		b.Annotations = make(map[string]string, len(value.(map[string]interface{})))
		for valuek, valuev := range value.(map[string]interface{}) {
			b.Annotations[valuek] = valuev.(string)
		}
	case "const":
		b.Const = schema.StringPtr(value.(string))
	case "default":
		b.Default = schema.StringPtr(value.(string))
	case "deprecated":
		b.Deprecated = value.(bool)
	case "description":
		b.Description = value.(string)
	case "enum":
		b.Enum = make([]string, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.Enum[valuek] = valuev.(string)
		}
	case "examples":
		b.Examples = value.([]interface{})
	case "format":
		b.Format = value.(string)
	case "id":
		b.ID = value.(string)
	case "max_length":
		b.MaxLength = schema.IntegerPtr(value.(int64))
	case "min_length":
		b.MinLength = value.(int64)
	case "nullable":
		b.Nullable = value.(bool)
	case "pattern":
		b.Pattern = formats.RegexPtr(value.(regexp.Regexp))
	case "read_only":
		b.ReadOnly = value.(bool)
	case "title":
		b.Title = value.(string)
	case "write_only":
		b.WriteOnly = value.(bool)
	}
	return nil
}

func (i StringInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
