// Code generated by Basil. DO NOT EDIT.

package blocks

import (
	"fmt"

	"github.com/opsidian/conflow/basil/schema"
	"github.com/opsidian/conflow/conflow"
)

// PrintInterpreter is the basil interpreter for the Print block
type PrintInterpreter struct {
	s schema.Schema
}

func (i PrintInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Metadata: schema.Metadata{
				Description: "It will write a string to the standard output",
			},
			Name: "Print",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
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
			Required: []string{"value"},
		}
	}
	return i.s
}

// Create creates a new Print block
func (i PrintInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &Print{
		id:     id,
		stdout: blockCtx.Stdout(),
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i PrintInterpreter) ValueParamName() conflow.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i PrintInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Print
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i PrintInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*Print).id
	case "value":
		return b.(*Print).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Print", name))
	}
}

func (i PrintInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Print)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i PrintInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
