package common

import (
	"fmt"

	"github.com/opsidian/basil/basil"
)

// Print will write a string to the standard output
//go:generate basil generate
type Print struct {
	id    basil.ID    `basil:"id"`
	value interface{} `basil:"value,required"`
}

func (p *Print) Main(ctx basil.BlockContext) error {
	fmt.Print(p.value)
	return nil
}
