// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/conflowio/conflow/src/conflow"
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{"block.conflow.io/eval_stage": "parse", "block.conflow.io/type": "directive"},
			ID:          "github.com/conflowio/conflow/src/directives.Object",
		},
		FieldNames:        map[string]string{"$id": "ID", "annotations": "Annotations", "const": "Const", "default": "Default", "dependentRequired": "DependentRequired", "deprecated": "Deprecated", "description": "Description", "enum": "Enum", "examples": "Examples", "maxProperties": "MaxProperties", "minProperties": "MinProperties", "readOnly": "ReadOnly", "title": "Title", "writeOnly": "WriteOnly"},
		JSONPropertyNames: map[string]string{"dependent_required": "dependentRequired", "id": "$id", "max_properties": "maxProperties", "min_properties": "minProperties", "property": "Parameters", "read_only": "readOnly", "write_only": "writeOnly"},
		Name:              "Object",
		Parameters: map[string]schema.Schema{
			"annotations": &schema.Map{
				AdditionalProperties: &schema.String{},
			},
			"const": &schema.Map{
				AdditionalProperties: &schema.Untyped{},
			},
			"default": &schema.Map{
				AdditionalProperties: &schema.Untyped{},
			},
			"dependent_required": &schema.Map{
				AdditionalProperties: &schema.Array{
					Items: &schema.String{},
				},
			},
			"deprecated":  &schema.Boolean{},
			"description": &schema.String{},
			"enum": &schema.Array{
				Items: &schema.Map{
					AdditionalProperties: &schema.Untyped{},
				},
			},
			"examples": &schema.Array{
				Items: &schema.Untyped{},
			},
			"id": &schema.String{},
			"max_properties": &schema.Integer{
				Nullable: true,
			},
			"min_properties": &schema.Integer{},
			"property": &schema.Map{
				AdditionalProperties: &schema.Reference{
					Ref: "github.com/conflowio/conflow/src/schema.Schema",
				},
			},
			"read_only":  &schema.Boolean{},
			"title":      &schema.String{},
			"write_only": &schema.Boolean{},
		},
	})
}

// ObjectInterpreter is the Conflow interpreter for the Object block
type ObjectInterpreter struct {
}

func (i ObjectInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/directives.Object")
	return s
}

// Create creates a new Object block
func (i ObjectInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Object{}
	b.Parameters = map[string]schema.Schema{}
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i ObjectInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i ObjectInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Object
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i ObjectInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "annotations":
		return b.(*Object).Annotations
	case "const":
		return b.(*Object).Const
	case "default":
		return b.(*Object).Default
	case "dependent_required":
		return b.(*Object).DependentRequired
	case "deprecated":
		return b.(*Object).Deprecated
	case "description":
		return b.(*Object).Description
	case "enum":
		return b.(*Object).Enum
	case "examples":
		return b.(*Object).Examples
	case "id":
		return b.(*Object).ID
	case "max_properties":
		return b.(*Object).MaxProperties
	case "min_properties":
		return b.(*Object).MinProperties
	case "read_only":
		return b.(*Object).ReadOnly
	case "title":
		return b.(*Object).Title
	case "write_only":
		return b.(*Object).WriteOnly
	default:
		panic(fmt.Errorf("unexpected parameter %q in Object", name))
	}
}

func (i ObjectInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Object)
	switch name {
	case "annotations":
		b.Annotations = make(map[string]string, len(value.(map[string]interface{})))
		for valuek, valuev := range value.(map[string]interface{}) {
			b.Annotations[valuek] = valuev.(string)
		}
	case "const":
		b.Const = value.(map[string]interface{})
	case "default":
		b.Default = value.(map[string]interface{})
	case "dependent_required":
		b.DependentRequired = make(map[string][]string, len(value.(map[string]interface{})))
		for valuek, valuev := range value.(map[string]interface{}) {
			b.DependentRequired[valuek] = make([]string, len(valuev.([]interface{})))
			for valuevk, valuevv := range valuev.([]interface{}) {
				b.DependentRequired[valuek][valuevk] = valuevv.(string)
			}
		}
	case "deprecated":
		b.Deprecated = value.(bool)
	case "description":
		b.Description = value.(string)
	case "enum":
		b.Enum = make([]map[string]interface{}, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.Enum[valuek] = valuev.(map[string]interface{})
		}
	case "examples":
		b.Examples = value.([]interface{})
	case "id":
		b.ID = value.(string)
	case "max_properties":
		b.MaxProperties = schema.IntegerPtr(value.(int64))
	case "min_properties":
		b.MinProperties = value.(int64)
	case "read_only":
		b.ReadOnly = value.(bool)
	case "title":
		b.Title = value.(string)
	case "write_only":
		b.WriteOnly = value.(bool)
	}
	return nil
}

func (i ObjectInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	b := block.(*Object)
	switch name {
	case "property":
		b.Parameters[key] = value.(schema.Schema)
	}
	return nil
}
