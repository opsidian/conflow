// Code generated by Basil. DO NOT EDIT.

package main

import (
	"fmt"

	"github.com/opsidian/conflow/basil/schema"
	"github.com/opsidian/conflow/conflow"
)

// BenchmarkRunInterpreter is the basil interpreter for the BenchmarkRun block
type BenchmarkRunInterpreter struct {
	s schema.Schema
}

func (i BenchmarkRunInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "BenchmarkRun",
			Properties: map[string]schema.Schema{
				"cnt": &schema.Integer{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"eval_stage": "close"},
						ReadOnly:    true,
					},
				},
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
			},
		}
	}
	return i.s
}

// Create creates a new BenchmarkRun block
func (i BenchmarkRunInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &BenchmarkRun{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BenchmarkRunInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BenchmarkRunInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *BenchmarkRun
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BenchmarkRunInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "cnt":
		return b.(*BenchmarkRun).cnt
	case "id":
		return b.(*BenchmarkRun).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in BenchmarkRun", name))
	}
}

func (i BenchmarkRunInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}

func (i BenchmarkRunInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	return nil
}
