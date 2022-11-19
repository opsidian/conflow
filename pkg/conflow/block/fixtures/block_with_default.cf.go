// Code generated by Conflow. DO NOT EDIT.

package fixtures

import (
	"fmt"

	"github.com/conflowio/conflow/pkg/conflow"
	"github.com/conflowio/conflow/pkg/conflow/annotations"
	"github.com/conflowio/conflow/pkg/schema"
)

func init() {
	schema.Register(&schema.Object{
		Metadata: schema.Metadata{
			Annotations: map[string]string{
				annotations.Type: "configuration",
			},
			ID: "github.com/conflowio/conflow/pkg/conflow/block/fixtures.BlockWithDefault",
		},
		ParameterNames: map[string]string{"IDField": "id_field", "Value": "value"},
		Properties: map[string]schema.Schema{
			"IDField": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
			"Value": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.Value: "true",
					},
				},
				Default: schema.Pointer("foo"),
			},
		},
	})
}

// NewBlockWithDefaultWithDefaults creates a new BlockWithDefault instance with default values
func NewBlockWithDefaultWithDefaults() *BlockWithDefault {
	b := &BlockWithDefault{}
	b.Value = "foo"
	return b
}

// BlockWithDefaultInterpreter is the Conflow interpreter for the BlockWithDefault block
type BlockWithDefaultInterpreter struct {
}

func (i BlockWithDefaultInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/pkg/conflow/block/fixtures.BlockWithDefault")
	return s
}

// Create creates a new BlockWithDefault block
func (i BlockWithDefaultInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := NewBlockWithDefaultWithDefaults()
	b.IDField = id
	return b
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
		b.Value = schema.Value[string](value)
	}
	return nil
}

func (i BlockWithDefaultInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
