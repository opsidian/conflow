package generator

type InterpreterTemplateParams struct {
	Package                string
	Name                   string
	FuncName               string
	Arguments              []*Argument
	Results                []*Argument
	ResultType             string
	NodeValueFunctionNames map[string]string
}

const interpreterTemplate = `
// Code generated by Basil. DO NOT EDIT.
package {{.Package}}

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/function"
	"github.com/opsidian/basil/util"
	"github.com/opsidian/parsley/parsley"
)

{{ $root := .}}

// {{.Name}}Interpreter is an AST node interpreter for the {{.FuncName}} function
type {{.Name}}Interpreter struct {}

// StaticCheck runs a static analysis on the function parameters
func (i {{.Name}}Interpreter) StaticCheck(ctx interface{}, node basil.FunctionNode) (string, parsley.Error) {
	if len(node.ArgumentNodes()) != {{ len .Arguments }} {
		return "", parsley.NewError(node.Pos(), fmt.Errorf("%s expects {{ len .Arguments }} arguments", node.Name()))
	}

	{{ if .Arguments }}
	arguments := node.ArgumentNodes()
	{{ range $i, $arg := .Arguments }}
	if err := util.CheckNodeType(arguments[{{$i}}], "{{.Type}}"); err != nil {
		return "", err
	}
	{{ end }}

	{{ end -}}
	return "{{.ResultType}}", nil
}

// Eval returns with the result of the function
func (i {{.Name}}Interpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	{{- if .Arguments }}
	arguments := node.ArgumentNodes()
	{{ range $i, $arg := .Arguments }}
	arg{{$i}}, err := util.{{index $root.NodeValueFunctionNames .Type}}(arguments[{{$i}}], ctx)
	if err != nil {
		return nil, err
	}
	{{ end }}

	{{ end -}}
	{{ if eq (len .Results) 1 }}
	return {{.FuncName}}(
		{{range $i, $arg := .Arguments -}}
		arg{{$i}},
		{{ end -}}
	), nil
	{{ else }}
	res, resErr := {{.FuncName}}(
		{{- range $i, $arg := .Arguments -}}
		arg{{$i}},
		{{- end -}}
	)
	if resErr != nil {
		if funcErr, ok := resErr.(*function.Error); ok {
			return nil, parsley.NewError(arguments[funcErr.ArgIndex].Pos(), funcErr.Err)
		}
		return nil, parsley.NewError(node.Pos(), resErr)
	}

	return res, nil
	{{ end }}
}
`
