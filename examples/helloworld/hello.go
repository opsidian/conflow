package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/opsidian/basil/basil"
)

// Hello is capable to print some greetings
//go:generate basil generate
type Hello struct {
	id       basil.ID   `basil:"id"`
	to       string     `basil:"required"`
	greeting string     `basil:"out"`
	r        *rand.Rand `basil:"ignore"`
}

// Init will initialise the random generator
func (h *Hello) Init(ctx basil.BlockContext) (bool, error) {
	h.r = rand.New(rand.NewSource(time.Now().Unix()))
	return true, nil
}

// Main will generate a random greeting
func (h *Hello) Main(ctx basil.BlockContext) error {
	greetings := []string{"Hello", "Hi", "Hey", "Yo", "Sup"}

	h.greeting = fmt.Sprintf("%s %s!", greetings[h.r.Intn(len(greetings))], h.to)

	return nil
}
