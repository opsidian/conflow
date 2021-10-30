// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package generator

import "github.com/opsidian/conflow/basil/schema"

type InterpreterTemplateParams struct {
	Package          string
	Name             string
	FuncNameSelector string
	FuncName         string
	Schema           schema.Schema
	Imports          map[string]string
	ReturnsError     bool
}

const interpreterHeaderTemplate = `
{{ $root := . -}}

// Code generated by Basil. DO NOT EDIT.

package {{ .Package }}

import (
	{{ range sortedImportKeys .Imports -}}
	{{ if ne . "." -}}
	{{ if ne (last (index $root.Imports .)) . }}{{ . }} {{ end }}{{ printf "%q" (index $root.Imports .) }}
	{{ end -}}
	{{ end -}}
)
`

const interpreterTemplate = `
{{ $root := . -}}

// {{ .Name }}Interpreter is the basil interpreter for the {{ .FuncName }} function
type {{ .Name }}Interpreter struct {
	s schema.Schema
}

func (i {{ .Name }}Interpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = {{ .Schema.GoString }}
	}
	return i.s
}

// Eval returns with the result of the function
func (i {{ .Name }}Interpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	{{ if .Schema.GetParameters -}}
	parameters := i.Schema().(*schema.Function).GetParameters()
	arguments := node.ArgumentNodes()

	{{ range $i, $property := .Schema.GetParameters -}}
	arg{{ $i }}, evalErr := parsley.EvaluateNode(ctx, arguments[{{ $i }}])
	if evalErr != nil {
		return nil, evalErr
	}
	if err := parameters[{{ $i }}].Schema.ValidateValue(arg{{ $i }}); err != nil {
		return nil, parsley.NewError(arguments[{{ $i }}].Pos(), err)
	}
	var {{ assignValue $property.Schema (printf "arg%d" $i) (printf "val%d" $i) }}

	{{ end -}}

	{{ end -}}
	{{ if .ReturnsError -}}
	res, resErr := {{ .FuncNameSelector }}{{ .FuncName }}(
		{{- range $i, $property := .Schema.GetParameters -}}
		val{{ $i }},
		{{- end -}}
	)
	if resErr != nil {
		if funcErr, ok := resErr.(*function.Error); ok {
			return nil, parsley.NewError(arguments[funcErr.ArgIndex].Pos(), funcErr.Err)
		}
		return nil, parsley.NewError(node.Pos(), resErr)
	}

	return res, nil
	{{ else -}}
	return {{ .FuncNameSelector }}{{ .FuncName }}(
		{{- range $i, $property := .Schema.GetParameters -}}
		val{{ $i }},
		{{- end -}}
	), nil
	{{ end -}}
}
`
