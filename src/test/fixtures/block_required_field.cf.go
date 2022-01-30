// Code generated by Conflow. DO NOT EDIT.

package fixtures

import (
	"fmt"
	"github.com/conflowio/conflow/src/conflow"
	"github.com/conflowio/conflow/src/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{"block.conflow.io/type": "configuration"},
			ID:          "github.com/conflowio/conflow/src/test/fixtures.BlockRequiredField",
		},
		JSONPropertyNames: map[string]string{"id_field": "IDField"},
		Name:              "BlockRequiredField",
		Parameters: map[string]schema.Schema{
			"id_field": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{"block.conflow.io/id": "true"},
					ReadOnly:    true,
				},
				Format: "conflow.ID",
			},
			"required": &schema.Untyped{},
		},
		Required: []string{"required"},
	})
}

// BlockRequiredFieldInterpreter is the Conflow interpreter for the BlockRequiredField block
type BlockRequiredFieldInterpreter struct {
}

func (i BlockRequiredFieldInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/test/fixtures.BlockRequiredField")
	return s
}

// Create creates a new BlockRequiredField block
func (i BlockRequiredFieldInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &BlockRequiredField{
		IDField: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockRequiredFieldInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockRequiredFieldInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *BlockRequiredField
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockRequiredFieldInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id_field":
		return b.(*BlockRequiredField).IDField
	case "required":
		return b.(*BlockRequiredField).required
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockRequiredField", name))
	}
}

func (i BlockRequiredFieldInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*BlockRequiredField)
	switch name {
	case "required":
		b.required = value
	}
	return nil
}

func (i BlockRequiredFieldInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
