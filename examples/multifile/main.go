package main

import (
	"path/filepath"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/block"
	"github.com/opsidian/basil/examples/common"
	"github.com/opsidian/basil/parser"
	"github.com/opsidian/basil/util"
)

//go:generate basil generate
type Main struct {
	id basil.ID `basil:"id"`
}

func (m *Main) ParseContextOverride() basil.ParseContextOverride {
	return basil.ParseContextOverride{
		BlockTransformerRegistry: block.InterpreterRegistry{
			"print":   common.PrintInterpreter{},
			"println": common.PrintlnInterpreter{},
		},
	}
}

func main() {
	ctx, cancel := util.CreateDefaultContext()
	defer cancel()

	parseCtx := basil.NewParseContext(basil.NewIDRegistry(8, 16))

	p := parser.NewMain("main", MainInterpreter{})

	paths, _ := filepath.Glob("*.basil")

	if err := p.ParseFiles(
		parseCtx,
		paths...,
	); err != nil {
		panic(err)
	}

	common.Main(ctx, parseCtx)
}
