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
			ID: "github.com/conflowio/conflow/src/test/fixtures.BlockNoFields",
		},
		JSONPropertyNames: map[string]string{"id_field": "IDField"},
		Parameters: map[string]schema.Schema{
			"id_field": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
		},
	})
}

// BlockNoFieldsInterpreter is the Conflow interpreter for the BlockNoFields block
type BlockNoFieldsInterpreter struct {
}

func (i BlockNoFieldsInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/test/fixtures.BlockNoFields")
	return s
}

// Create creates a new BlockNoFields block
func (i BlockNoFieldsInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &BlockNoFields{}
	b.IDField = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BlockNoFieldsInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockNoFieldsInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *BlockNoFields
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockNoFieldsInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id_field":
		return b.(*BlockNoFields).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockNoFields", name))
	}
}

func (i BlockNoFieldsInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}

func (i BlockNoFieldsInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
