// Code generated by Conflow. DO NOT EDIT.

package blocks

import (
	"fmt"
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
	"io"
)

// GzipInterpreter is the conflow interpreter for the Gzip block
type GzipInterpreter struct {
	s schema.Schema
}

func (i GzipInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Metadata: schema.Metadata{
				Annotations: map[string]string{"block.conflow.io/type": "task"},
			},
			Name: "Gzip",
			Parameters: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"block.conflow.io/id": "true"},
						ReadOnly:    true,
					},
					Format: "conflow.ID",
				},
				"in": &schema.ByteStream{},
				"out": &schema.Reference{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"block.conflow.io/eval_stage": "init", "block.conflow.io/generated": "true"},
					},
					Nullable: true,
					Ref:      "http://conflow.schema/github.com/conflowio/conflow/blocks.Stream",
				},
			},
			Required: []string{"in", "out"},
		}
	}
	return i.s
}

// Create creates a new Gzip block
func (i GzipInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &Gzip{
		id:             id,
		blockPublisher: blockCtx.BlockPublisher(),
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i GzipInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i GzipInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Gzip
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i GzipInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "id":
		return b.(*Gzip).id
	case "in":
		return b.(*Gzip).in
	default:
		panic(fmt.Errorf("unexpected parameter %q in Gzip", name))
	}
}

func (i GzipInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Gzip)
	switch name {
	case "in":
		b.in = value.(io.ReadCloser)
	}
	return nil
}

func (i GzipInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Gzip)
	switch name {
	case "out":
		b.out = value.(*Stream)
	}
	return nil
}
