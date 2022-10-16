// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package openapi

import (
	"github.com/conflowio/conflow/src/conflow"
	"github.com/conflowio/conflow/src/conflow/block"
	"github.com/conflowio/conflow/src/schema"
)

// @block "configuration"
type PathItem struct {
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Get         *Operation `json:"get,omitempty"`
	Put         *Operation `json:"put,omitempty"`
	Post        *Operation `json:"post,omitempty"`
	Delete      *Operation `json:"delete,omitempty"`
	Options     *Operation `json:"options,omitempty"`
	Head        *Operation `json:"head,omitempty"`
	Patch       *Operation `json:"patch,omitempty"`
	Trace       *Operation `json:"trace,omitempty"`
	// @name "server"
	Servers []*Server `json:"servers,omitempty"`
	// @name "parameter"
	Parameters []*Parameter `json:"parameters,omitempty"`
}

func (p *PathItem) ParseContextOverride() conflow.ParseContextOverride {
	return conflow.ParseContextOverride{
		BlockTransformerRegistry: block.InterpreterRegistry{
			"get":       OperationInterpreter{},
			"put":       OperationInterpreter{},
			"post":      OperationInterpreter{},
			"delete":    OperationInterpreter{},
			"options":   OperationInterpreter{},
			"head":      OperationInterpreter{},
			"patch":     OperationInterpreter{},
			"trace":     OperationInterpreter{},
			"server":    ServerInterpreter{},
			"parameter": ParameterInterpreter{},
		},
	}
}

func (p *PathItem) Validate(ctx *schema.Context) error {
	return schema.ValidateAll(
		ctx,
		schema.Validate("get", p.Get),
		schema.Validate("put", p.Put),
		schema.Validate("post", p.Post),
		schema.Validate("delete", p.Delete),
		schema.Validate("options", p.Options),
		schema.Validate("head", p.Head),
		schema.Validate("patch", p.Patch),
		schema.Validate("trace", p.Trace),
		schema.ValidateArray("servers", p.Servers),
		schema.ValidateArray("parameters", p.Parameters),
	)
}
