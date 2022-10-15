// Code generated by Conflow. DO NOT EDIT.

package fixtures

import (
	"fmt"
	"github.com/conflowio/conflow/src/conflow"
	"github.com/conflowio/conflow/src/conflow/annotations"
	"github.com/conflowio/conflow/src/schema"
	"time"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{
				annotations.Type: "configuration",
			},
			ID: "github.com/conflowio/conflow/src/test/fixtures.Block",
		},
		JSONPropertyNames: map[string]string{"field_array": "FieldArray", "field_bool": "FieldBool", "field_float": "FieldFloat", "field_identifier": "FieldIdentifier", "field_integer": "FieldInteger", "field_interface": "FieldInterface", "field_map": "FieldMap", "field_number": "FieldNumber", "field_string": "FieldString", "field_string_array": "FieldStringArray", "field_time": "FieldTime", "field_time_duration": "FieldTimeDuration", "id_field": "IDField"},
		Parameters: map[string]schema.Schema{
			"field_array": &schema.Array{
				Items: &schema.Untyped{},
			},
			"field_bool":  &schema.Boolean{},
			"field_float": &schema.Number{},
			"field_identifier": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.EvalStage: "close",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
			"field_integer":   &schema.Integer{},
			"field_interface": &schema.Untyped{},
			"field_map": &schema.Map{
				AdditionalProperties: &schema.Untyped{},
			},
			"field_number": &schema.Untyped{
				Types: []string{"integer", "number"},
			},
			"field_string": &schema.String{},
			"field_string_array": &schema.Array{
				Items: &schema.String{},
			},
			"field_time": &schema.String{
				Format: "date-time",
			},
			"field_time_duration": &schema.String{
				Format: "duration-go",
			},
			"id_field": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
		},
	})
}

// BlockInterpreter is the Conflow interpreter for the Block block
type BlockInterpreter struct {
}

func (i BlockInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/test/fixtures.Block")
	return s
}

// Create creates a new Block block
func (i BlockInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &Block{}
	b.IDField = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Block
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "field_array":
		return b.(*Block).FieldArray
	case "field_bool":
		return b.(*Block).FieldBool
	case "field_float":
		return b.(*Block).FieldFloat
	case "field_identifier":
		return b.(*Block).FieldIdentifier
	case "field_integer":
		return b.(*Block).FieldInteger
	case "field_interface":
		return b.(*Block).FieldInterface
	case "field_map":
		return b.(*Block).FieldMap
	case "field_number":
		return b.(*Block).FieldNumber
	case "field_string":
		return b.(*Block).FieldString
	case "field_string_array":
		return b.(*Block).FieldStringArray
	case "field_time":
		return b.(*Block).FieldTime
	case "field_time_duration":
		return b.(*Block).FieldTimeDuration
	case "id_field":
		return b.(*Block).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in Block", name))
	}
}

func (i BlockInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Block)
	switch name {
	case "field_array":
		b.FieldArray = value.([]interface{})
	case "field_bool":
		b.FieldBool = value.(bool)
	case "field_float":
		b.FieldFloat = value.(float64)
	case "field_integer":
		b.FieldInteger = value.(int64)
	case "field_interface":
		b.FieldInterface = value
	case "field_map":
		b.FieldMap = value.(map[string]interface{})
	case "field_number":
		b.FieldNumber = value
	case "field_string":
		b.FieldString = value.(string)
	case "field_string_array":
		b.FieldStringArray = make([]string, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.FieldStringArray[valuek] = valuev.(string)
		}
	case "field_time":
		b.FieldTime = value.(time.Time)
	case "field_time_duration":
		b.FieldTimeDuration = value.(time.Duration)
	}
	return nil
}

func (i BlockInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
