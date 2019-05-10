package generator

type InterpreterTemplateParams struct {
	Package                string
	Type                   string
	Name                   string
	HasForeignID           bool
	Stages                 []string
	ValueField             *Field
	Params                 []*Field
	Blocks                 []*Field
	RequiredParams         []string
	IDField                *Field
	NodeValueFunctionNames map[string]string
}

const interpreterTemplate = `
// Code generated by Basil. DO NOT EDIT.
package {{.Package}}

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/util"
	"github.com/opsidian/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

{{ $root := .}}

type {{.Name}}Interpreter struct {}

// Create creates a new {{.Name}} block
func (i {{.Name}}Interpreter) Create(ctx *basil.EvalContext, node basil.BlockNode) basil.Block {
	return &{{.Name}}{
		{{.IDField.Name}}: node.ID(),
	}
}

// Params returns with the list of valid parameters
func (i {{.Name}}Interpreter) Params() map[basil.ID]block.ParameterDescriptor {
	{{ if .Params -}}
	return map[basil.ID]block.ParameterDescriptor{
		{{ range .Params -}}
		"{{.ParamName}}": { Type: "{{.Type}}", IsRequired: {{.IsRequired}}, IsOutput: {{.IsOutput}}},
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
func (i {{.Name}}Interpreter) ParseContext(parentCtx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *{{.Name}}
	if b, ok := basil.Block(nilBlock).(basil.ParseContextAware); ok {
		return b.ParseContext(parentCtx)
	}

	return parentCtx
}

func (i {{.Name}}Interpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	{{ range .Params }}
	case "{{.ParamName}}":
		return b.(*{{$root.Name}}).{{.Name}}
	{{ end -}}
	default:
		panic(fmt.Errorf("unexpected parameter %q in {{.Name}}", name))
	}
}

func (i {{.Name}}Interpreter) SetParam(ctx *basil.EvalContext, b basil.Block, name basil.ID, node basil.BlockParamNode) parsley.Error {
	switch name {
	{{ range .Params -}}
	{{ if and (not .IsID) (not .IsOutput) -}}
	case "{{.ParamName}}":
		var err parsley.Error
		b.(*{{$root.Name}}).{{.Name}}, err = variable.{{index $root.NodeValueFunctionNames .Type}}(node, ctx)
		return err
	{{ end -}}
	{{ end -}}
	}

	return nil
}

func (i {{.Name}}Interpreter) SetBlock(ctx *basil.EvalContext, b basil.Block, name basil.ID, value interface{}) parsley.Error {
	{{ if .Blocks -}}
	switch name {
	{{ range .Blocks -}}
	case "{{.ParamName}}":
		b.(*{{$root.Name}}).{{.Name}} = append(b.(*{{$root.Name}}).{{.Name}}, value.({{trimPrefix .Type "[]" }}))
	{{ end -}}
	}
	{{ end -}}
	return nil
}

`
