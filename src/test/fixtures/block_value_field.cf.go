// Code generated by Conflow. DO NOT EDIT.

package fixtures

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
				annotations.Type: "configuration",
			},
			ID: "github.com/conflowio/conflow/src/test/fixtures.BlockValueField",
		},
		JSONPropertyNames: map[string]string{"id_field": "IDField"},
		ParameterNames:    map[string]string{"IDField": "id_field"},
		Properties: map[string]schema.Schema{
			"IDField": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
			"value": &schema.Any{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.Value: "true",
					},
				},
			},
		},
	})
}

// BlockValueFieldInterpreter is the Conflow interpreter for the BlockValueField block
type BlockValueFieldInterpreter struct {
}

func (i BlockValueFieldInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/test/fixtures.BlockValueField")
	return s
}

// Create creates a new BlockValueField block
func (i BlockValueFieldInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &BlockValueField{}
	b.IDField = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockValueFieldInterpreter) ValueParamName() conflow.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i BlockValueFieldInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *BlockValueField
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockValueFieldInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id_field":
		return b.(*BlockValueField).IDField
	case "value":
		return b.(*BlockValueField).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockValueField", name))
	}
}

func (i BlockValueFieldInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*BlockValueField)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i BlockValueFieldInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
