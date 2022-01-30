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
			Annotations: map[string]string{"block.conflow.io/eval_stage": "ignore", "block.conflow.io/type": "directive"},
			ID:          "github.com/conflowio/conflow/src/directives.Doc",
		},
		Name: "Doc",
		Parameters: map[string]schema.Schema{
			"description": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{"block.conflow.io/value": "true"},
				},
			},
			"id": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{"block.conflow.io/id": "true"},
					ReadOnly:    true,
				},
				Format: "conflow.ID",
			},
		},
		Required: []string{"description"},
	})
}

// DocInterpreter is the Conflow interpreter for the Doc block
type DocInterpreter struct {
}

func (i DocInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/directives.Doc")
	return s
}

// Create creates a new Doc block
func (i DocInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &Doc{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i DocInterpreter) ValueParamName() conflow.ID {
	return "description"
}

// ParseContext returns with the parse context for the block
func (i DocInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Doc
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i DocInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "description":
		return b.(*Doc).description
	case "id":
		return b.(*Doc).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Doc", name))
	}
}

func (i DocInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Doc)
	switch name {
	case "description":
		b.description = value.(string)
	}
	return nil
}

func (i DocInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
