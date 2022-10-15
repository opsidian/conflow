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
			ID: "github.com/conflowio/conflow/src/test/fixtures.BlockGeneratorResult",
		},
		Name: "BlockGeneratorResult",
		Parameters: map[string]schema.Schema{
			"id": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
			"value": &schema.Untyped{},
		},
	})
}

// BlockGeneratorResultInterpreter is the Conflow interpreter for the BlockGeneratorResult block
type BlockGeneratorResultInterpreter struct {
}

func (i BlockGeneratorResultInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/test/fixtures.BlockGeneratorResult")
	return s
}

// Create creates a new BlockGeneratorResult block
func (i BlockGeneratorResultInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &BlockGeneratorResult{}
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockGeneratorResultInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockGeneratorResultInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *BlockGeneratorResult
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockGeneratorResultInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*BlockGeneratorResult).id
	case "value":
		return b.(*BlockGeneratorResult).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockGeneratorResult", name))
	}
}

func (i BlockGeneratorResultInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*BlockGeneratorResult)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i BlockGeneratorResultInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
