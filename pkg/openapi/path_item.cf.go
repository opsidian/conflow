// Code generated by Conflow. DO NOT EDIT.

package openapi

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
			ID: "github.com/conflowio/conflow/pkg/openapi.PathItem",
		},
		FieldNames:     map[string]string{"delete": "Delete", "description": "Description", "get": "Get", "head": "Head", "options": "Options", "parameters": "Parameters", "patch": "Patch", "post": "Post", "put": "Put", "servers": "Servers", "summary": "Summary", "trace": "Trace"},
		ParameterNames: map[string]string{"parameters": "parameter", "servers": "server"},
		Properties: map[string]schema.Schema{
			"delete": &schema.Reference{
				Nullable: true,
				Ref:      "github.com/conflowio/conflow/pkg/openapi.Operation",
			},
			"description": &schema.String{},
			"get": &schema.Reference{
				Nullable: true,
				Ref:      "github.com/conflowio/conflow/pkg/openapi.Operation",
			},
			"head": &schema.Reference{
				Nullable: true,
				Ref:      "github.com/conflowio/conflow/pkg/openapi.Operation",
			},
			"options": &schema.Reference{
				Nullable: true,
				Ref:      "github.com/conflowio/conflow/pkg/openapi.Operation",
			},
			"parameters": &schema.Array{
				Items: &schema.Reference{
					Nullable: true,
					Ref:      "github.com/conflowio/conflow/pkg/openapi.Parameter",
				},
			},
			"patch": &schema.Reference{
				Nullable: true,
				Ref:      "github.com/conflowio/conflow/pkg/openapi.Operation",
			},
			"post": &schema.Reference{
				Nullable: true,
				Ref:      "github.com/conflowio/conflow/pkg/openapi.Operation",
			},
			"put": &schema.Reference{
				Nullable: true,
				Ref:      "github.com/conflowio/conflow/pkg/openapi.Operation",
			},
			"servers": &schema.Array{
				Items: &schema.Reference{
					Nullable: true,
					Ref:      "github.com/conflowio/conflow/pkg/openapi.Server",
				},
			},
			"summary": &schema.String{},
			"trace": &schema.Reference{
				Nullable: true,
				Ref:      "github.com/conflowio/conflow/pkg/openapi.Operation",
			},
		},
	})
}

// PathItemInterpreter is the Conflow interpreter for the PathItem block
type PathItemInterpreter struct {
}

func (i PathItemInterpreter) Schema() schema.Schema {
	s, _ := schema.Get("github.com/conflowio/conflow/pkg/openapi.PathItem")
	return s
}

// Create creates a new PathItem block
func (i PathItemInterpreter) CreateBlock(id conflow.ID, blockCtx *conflow.BlockContext) conflow.Block {
	b := &PathItem{}
	return b
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i PathItemInterpreter) ValueParamName() conflow.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i PathItemInterpreter) ParseContext(ctx *conflow.ParseContext) *conflow.ParseContext {
	var nilBlock *PathItem
	if b, ok := conflow.Block(nilBlock).(conflow.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i PathItemInterpreter) Param(b conflow.Block, name conflow.ID) interface{} {
	switch name {
	case "description":
		return b.(*PathItem).Description
	case "summary":
		return b.(*PathItem).Summary
	default:
		panic(fmt.Errorf("unexpected parameter %q in PathItem", name))
	}
}

func (i PathItemInterpreter) SetParam(block conflow.Block, name conflow.ID, value interface{}) error {
	b := block.(*PathItem)
	switch name {
	case "description":
		b.Description = value.(string)
	case "summary":
		b.Summary = value.(string)
	}
	return nil
}

func (i PathItemInterpreter) SetBlock(block conflow.Block, name conflow.ID, key string, value interface{}) error {
	b := block.(*PathItem)
	switch name {
	case "delete":
		b.Delete = value.(*Operation)
	case "get":
		b.Get = value.(*Operation)
	case "head":
		b.Head = value.(*Operation)
	case "options":
		b.Options = value.(*Operation)
	case "parameter":
		b.Parameters = append(b.Parameters, value.(*Parameter))
	case "patch":
		b.Patch = value.(*Operation)
	case "post":
		b.Post = value.(*Operation)
	case "put":
		b.Put = value.(*Operation)
	case "server":
		b.Servers = append(b.Servers, value.(*Server))
	case "trace":
		b.Trace = value.(*Operation)
	}
	return nil
}
