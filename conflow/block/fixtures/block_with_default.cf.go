// Code generated by Conflow. DO NOT EDIT.

package fixtures

import (
	"fmt"
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
)

// BlockWithDefaultInterpreter is the conflow interpreter for the BlockWithDefault block
type BlockWithDefaultInterpreter struct {
	s schema.Schema
}

func (i BlockWithDefaultInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "BlockWithDefault",
			Properties: map[string]schema.Schema{
				"id_field": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "conflow.ID",
				},
				"value": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"value": "true"},
					},
					Default: schema.StringPtr("foo"),
				},
			},
			PropertyNames: map[string]string{"id_field": "IDField", "value": "Value"},
		}
	}
	return i.s
}

// Create creates a new BlockWithDefault block
func (i BlockWithDefaultInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &BlockWithDefault{
		IDField: id,
		Value:   "foo",
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockWithDefaultInterpreter) ValueParamName() conflow.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i BlockWithDefaultInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *BlockWithDefault
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockWithDefaultInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id_field":
		return b.(*BlockWithDefault).IDField
	case "value":
		return b.(*BlockWithDefault).Value
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockWithDefault", name))
	}
}

func (i BlockWithDefaultInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*BlockWithDefault)
	switch name {
	case "value":
		b.Value = value.(string)
	}
	return nil
}

func (i BlockWithDefaultInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
