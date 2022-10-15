// Code generated by Conflow. DO NOT EDIT.

package directives

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
				annotations.Type: "directive",
			},
			ID: "github.com/conflowio/conflow/src/schema/directives.MultipleOf",
		},
		Name: "MultipleOf",
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
			"value": &schema.Untyped{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.Value: "true",
					},
				},
			},
		},
		Required: []string{"value"},
	})
}

// MultipleOfInterpreter is the Conflow interpreter for the MultipleOf block
type MultipleOfInterpreter struct {
}

func (i MultipleOfInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/src/schema/directives.MultipleOf")
	return s
}

// Create creates a new MultipleOf block
func (i MultipleOfInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &MultipleOf{}
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i MultipleOfInterpreter) ValueParamName() conflow.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i MultipleOfInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *MultipleOf
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i MultipleOfInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*MultipleOf).id
	case "value":
		return b.(*MultipleOf).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in MultipleOf", name))
	}
}

func (i MultipleOfInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*MultipleOf)
	switch name {
	case "value":
		b.value = value
	}
	return nil
}

func (i MultipleOfInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
