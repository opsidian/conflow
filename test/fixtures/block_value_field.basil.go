// Code generated by Basil. DO NOT EDIT.

package fixtures

import (
	"fmt"

	"github.com/opsidian/conflow/basil/schema"
	"github.com/opsidian/conflow/conflow"
)

// BlockValueFieldInterpreter is the basil interpreter for the BlockValueField block
type BlockValueFieldInterpreter struct {
	s schema.Schema
}

func (i BlockValueFieldInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "BlockValueField",
			Properties: map[string]schema.Schema{
				"id_field": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
				"value": &schema.Untyped{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"value": "true"},
					},
				},
			},
			PropertyNames: map[string]string{"id_field": "IDField"},
		}
	}
	return i.s
}

// Create creates a new BlockValueField block
func (i BlockValueFieldInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &BlockValueField{
		IDField: id,
	}
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

func (i BlockValueFieldInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
