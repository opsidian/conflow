// Code generated by Conflow. DO NOT EDIT.

package main

import (
	"fmt"
	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
)

// GlobInterpreter is the conflow interpreter for the Glob block
type GlobInterpreter struct {
	s schema.Schema
}

func (i GlobInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "Glob",
			Properties: map[string]schema.Schema{
				"exclude": &schema.Array{
					Items: &schema.String{},
				},
				"file": &schema.Reference{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"eval_stage": "init", "generated": "true"},
						Pointer:     true,
					},
					Ref: "http://conflow.schema/github.com/conflowio/conflow/examples/licensify.File",
				},
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "conflow.ID",
				},
				"include": &schema.Array{
					Items: &schema.String{},
				},
				"path": &schema.String{},
			},
			Required: []string{"path", "file"},
		}
	}
	return i.s
}

// Create creates a new Glob block
func (i GlobInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	return &Glob{
		id:             id,
		blockPublisher: blockCtx.BlockPublisher(),
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i GlobInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i GlobInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *Glob
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i GlobInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "exclude":
		return b.(*Glob).exclude
	case "id":
		return b.(*Glob).id
	case "include":
		return b.(*Glob).include
	case "path":
		return b.(*Glob).path
	default:
		panic(fmt.Errorf("unexpected parameter %q in Glob", name))
	}
}

func (i GlobInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Glob)
	switch name {
	case "exclude":
		b.exclude = make([]string, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.exclude[valuek] = valuev.(string)
		}
	case "include":
		b.include = make([]string, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.include[valuek] = valuev.(string)
		}
	case "path":
		b.path = value.(string)
	}
	return nil
}

func (i GlobInterpreter) SetBlock(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*Glob)
	switch name {
	case "file":
		b.file = value.(*File)
	}
	return nil
}
