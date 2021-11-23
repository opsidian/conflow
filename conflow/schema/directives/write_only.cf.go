// Code generated by Conflow. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
)

// WriteOnlyInterpreter is the conflow interpreter for the WriteOnly block
type WriteOnlyInterpreter struct {
	s schema.Schema
}

func (i WriteOnlyInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Metadata: schema.Metadata{
				Annotations: map[string]string{"block.conflow.io/type": "directive"},
			},
			Name: "WriteOnly",
			Parameters: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"block.conflow.io/id": "true"},
						ReadOnly:    true,
					},
					Format: "conflow.ID",
				},
			},
		}
	}
	return i.s
}

// Create creates a new WriteOnly block
func (i WriteOnlyInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &WriteOnly{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i WriteOnlyInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i WriteOnlyInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *WriteOnly
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i WriteOnlyInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*WriteOnly).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in WriteOnly", name))
	}
}

func (i WriteOnlyInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}

func (i WriteOnlyInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
