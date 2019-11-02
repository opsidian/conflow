// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package generator

type InterpreterTemplateParams struct {
	Package            string
	Type               string
	Name               string
	Fields             Fields
	IDField            *Field
	HasForeignID       bool
	ValueField         *Field
	ValueFunctionNames map[string]string
}

const interpreterTemplate = `
// Code generated by Basil. DO NOT EDIT.
package {{.Package}}

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/util"
	"github.com/opsidian/basil/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

{{ $root := .}}

type {{.Name}}Interpreter struct {}

// Create creates a new {{.Name}} block
func (i {{.Name}}Interpreter) CreateBlock(id basil.ID) basil.Block {
	return &{{.Name}}{
		{{.IDField.Name}}: id,
		{{ range (filterDefaults (filterParams .Fields)) -}}
		{{.Name}}: {{printf "%#v" .Default}},
		{{ end }}
	}
}

// Params returns with the list of valid parameters
func (i {{.Name}}Interpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	{{ if filterNonID (filterParams .Fields) -}}
	return map[basil.ID]basil.ParameterDescriptor{
		{{ range (filterNonID (filterParams .Fields)) -}}
		"{{.ParamName}}": {
			Type: "{{.Type}}",
			EvalStage: basil.EvalStages["{{.Stage}}"],
			IsRequired: {{.IsRequired}},
			IsOutput: {{.IsOutput}},
		},
		{{ end -}}
	}
	{{ else -}}
	return nil
	{{ end -}}
}

// Blocks returns with the list of valid blocks
func (i {{.Name}}Interpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	{{ if filterBlocks .Fields -}}
	return map[basil.ID]basil.BlockDescriptor{
		{{ range (filterBlocks .Fields) -}}
		"{{.ParamName}}": {
			EvalStage: basil.EvalStages["{{.Stage}}"],
			IsGenerated: {{.IsGenerated}},
			IsRequired: {{.IsRequired}},
			IsMany: {{.IsMany}},
		},
		{{ end -}}
	}
	{{ else -}}
	return nil
	{{ end -}}
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i {{.Name}}Interpreter) HasForeignID() bool {
	return {{.HasForeignID}}
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i {{.Name}}Interpreter) ValueParamName() basil.ID {
	return {{ if .ValueField }}"{{.ValueField.ParamName}}"{{ else }}""{{ end }}
}

// ParseContext returns with the parse context for the block
func (i {{.Name}}Interpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *{{.Name}}
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i {{.Name}}Interpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	{{ range filterParams .Fields -}}
	case "{{.ParamName}}":
		return b.(*{{$root.Name}}).{{.Name}}
	{{ end -}}
	default:
		panic(fmt.Errorf("unexpected parameter %q in {{.Name}}", name))
	}
}

func (i {{.Name}}Interpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	{{ if filterInputs (filterParams .Fields) -}}
	var err error
	b := block.(*{{$root.Name}})
	switch name {
	{{ range (filterInputs (filterParams .Fields)) -}}
	case "{{.ParamName}}":
		b.{{.Name}}, err = variable.{{index $root.ValueFunctionNames .Type}}(value)
	{{ end -}}
	}
	return err
	{{ else -}}
	return nil
	{{ end -}}
}

func (i {{.Name}}Interpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	{{ if filterInputs (filterBlocks .Fields) -}}
	b := block.(*{{$root.Name}})
	switch name {
	{{ range (filterInputs (filterBlocks .Fields)) -}}
	case "{{.ParamName}}":
		{{ if .IsMany -}}
		b.{{.Name}} = append(b.{{.Name}}, value.({{ .Type }}))
		{{ else -}}
		b.{{.Name}} = value.({{ .Type }})
		{{ end -}}
	{{ end -}}
	}
	{{ end -}}
	return nil
}

func (i {{.Name}}Interpreter) EvalStage() basil.EvalStage {
	var nilBlock *{{.Name}}
	if b, ok := basil.Block(nilBlock).(basil.EvalStageAware); ok {
		return b.EvalStage()
	}

	return basil.EvalStageUndefined
}

`
