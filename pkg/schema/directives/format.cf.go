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
			ID: "github.com/conflowio/conflow/pkg/schema/directives.Format",
		},
		Properties: map[string]schema.Schema{
			"id": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.ID: "true",
					},
					ReadOnly: true,
				},
				Format: "conflow.ID",
			},
			"value": &schema.String{
				Metadata: schema.Metadata{
					Annotations: map[string]string{
						annotations.Value: "true",
					},
				},
				Enum: []string{"date", "date-time", "duration", "email", "hostname", "idn-email", "idn-hostname", "ip", "ip-cidr", "ipv4", "ipv4-cidr", "ipv6", "ipv6-cidr", "iri", "iri-reference", "regex", "time", "uri", "uri-reference", "uri-template", "uuid"},
			},
		},
		Required: []string{"value"},
	})
}

// NewFormatWithDefaults creates a new Format instance with default values
func NewFormatWithDefaults() *Format {
	b := &Format{}
	return b
}

// FormatInterpreter is the Conflow interpreter for the Format block
type FormatInterpreter struct {
}

func (i FormatInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/pkg/schema/directives.Format")
	return s
}

// Create creates a new Format block
func (i FormatInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := NewFormatWithDefaults()
	b.id = id
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i FormatInterpreter) ValueParamName() conflow.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i FormatInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Format
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i FormatInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*Format).id
	case "value":
		return b.(*Format).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Format", name))
	}
}

func (i FormatInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Format)
	switch name {
	case "value":
		b.value = schema.Value[string](value)
	}
	return nil
}

func (i FormatInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	return nil
}
