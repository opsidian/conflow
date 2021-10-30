// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"

	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/schema"
)

// TitleInterpreter is the basil interpreter for the Title block
type TitleInterpreter struct {
	s schema.Schema
}

func (i TitleInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "Title",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
				"value": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"value": "true"},
					},
				},
			},
			Required: []string{"value"},
		}
	}
	return i.s
}

// Create creates a new Title block
func (i TitleInterpreter) CreateBlock(id basil.ID, blockCtx *basil.BlockContext) basil.Block {
	return &Title{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i TitleInterpreter) ValueParamName() basil.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i TitleInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Title
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i TitleInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Title).id
	case "value":
		return b.(*Title).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Title", name))
	}
}

func (i TitleInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Title)
	switch name {
	case "value":
		b.value = value.(string)
	}
	return nil
}

func (i TitleInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
