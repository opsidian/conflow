// Code generated by Conflow. DO NOT EDIT.

package directives

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
				annotations.Type: "directive",
			},
			ID: "github.com/conflowio/conflow/pkg/schema/directives.DependentRequired",
		},
		ParameterNames: map[string]string{"Value": "value"},
		Properties: map[string]schema.Schema{
			"Value": &schema.Map{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.Value: "true",
					},
				},
				AdditionalProperties: &schema.Array{
					Items: &schema.String{
						MinLength: 1,
					},
					MinItems:    1,
					UniqueItems: true,
				},
			},
			"id": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
		},
		Required: []string{"value"},
	})
}

// NewDependentRequiredWithDefaults creates a new DependentRequired instance with default values
func NewDependentRequiredWithDefaults() *DependentRequired {
	b := &DependentRequired{}
	return b
}

// DependentRequiredInterpreter is the Conflow interpreter for the DependentRequired block
type DependentRequiredInterpreter struct {
}

func (i DependentRequiredInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/pkg/schema/directives.DependentRequired")
	return s
}

// Create creates a new DependentRequired block
func (i DependentRequiredInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := NewDependentRequiredWithDefaults()
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i DependentRequiredInterpreter) ValueParamName() conflow.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i DependentRequiredInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *DependentRequired
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i DependentRequiredInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "value":
		return b.(*DependentRequired).Value
	case "id":
		return b.(*DependentRequired).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in DependentRequired", name))
	}
}

func (i DependentRequiredInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*DependentRequired)
	switch name {
	case "value":
		b.Value = make(map[string][]string, len(value.(map[string]interface{})))
		for valuek, valuev := range value.(map[string]interface{}) {
			b.Value[valuek] = make([]string, len(valuev.([]interface{})))
			for valuevk, valuevv := range valuev.([]interface{}) {
				b.Value[valuek][valuevk] = schema.Value[string](valuevv)
			}
		}
	}
	return nil
}

func (i DependentRequiredInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
