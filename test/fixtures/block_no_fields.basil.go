// Code generated by Basil. DO NOT EDIT.

package fixtures

import (
	"fmt"

	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/schema"
)

// BlockNoFieldsInterpreter is the basil interpreter for the BlockNoFields block
type BlockNoFieldsInterpreter struct {
	s schema.Schema
}

func (i BlockNoFieldsInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "BlockNoFields",
			Properties: map[string]schema.Schema{
				"id_field": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
			},
			PropertyNames: map[string]string{"id_field": "IDField"},
		}
	}
	return i.s
}

// Create creates a new BlockNoFields block
func (i BlockNoFieldsInterpreter) CreateBlock(id basil.ID, blockCtx *basil.BlockContext) basil.Block {
	return &BlockNoFields{
		IDField: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockNoFieldsInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockNoFieldsInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BlockNoFields
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockNoFieldsInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id_field":
		return b.(*BlockNoFields).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockNoFields", name))
	}
}

func (i BlockNoFieldsInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i BlockNoFieldsInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
