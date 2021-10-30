// Code generated by Basil. DO NOT EDIT.

package main

import (
	"fmt"
	"time"

	"github.com/opsidian/conflow/basil"
	"github.com/opsidian/conflow/basil/schema"
)

// BenchmarkInterpreter is the basil interpreter for the Benchmark block
type BenchmarkInterpreter struct {
	s schema.Schema
}

func (i BenchmarkInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "Benchmark",
			Properties: map[string]schema.Schema{
				"counter": &schema.Integer{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"eval_stage": "close"},
						ReadOnly:    true,
					},
				},
				"duration": &schema.TimeDuration{},
				"elapsed":  &schema.TimeDuration{},
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
				"run": &schema.Reference{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"eval_stage": "init", "generated": "true"},
						Pointer:     true,
					},
					Ref: "http://basil.schema/github.com/opsidian/conflow/examples/benchmark.BenchmarkRun",
				},
			},
			Required: []string{"duration", "run"},
		}
	}
	return i.s
}

// Create creates a new Benchmark block
func (i BenchmarkInterpreter) CreateBlock(id basil.ID, blockCtx *basil.BlockContext) basil.Block {
	return &Benchmark{
		id:             id,
		blockPublisher: blockCtx.BlockPublisher(),
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i BenchmarkInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BenchmarkInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Benchmark
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BenchmarkInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "counter":
		return b.(*Benchmark).counter
	case "duration":
		return b.(*Benchmark).duration
	case "elapsed":
		return b.(*Benchmark).elapsed
	case "id":
		return b.(*Benchmark).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Benchmark", name))
	}
}

func (i BenchmarkInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Benchmark)
	switch name {
	case "duration":
		b.duration = value.(time.Duration)
	case "elapsed":
		b.elapsed = value.(time.Duration)
	}
	return nil
}

func (i BenchmarkInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Benchmark)
	switch name {
	case "run":
		b.run = value.(*BenchmarkRun)
	}
	return nil
}
