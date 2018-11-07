package block_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opsidian/basil/block"
	"github.com/opsidian/basil/block/fixtures"
	"github.com/opsidian/basil/parser"
	"github.com/opsidian/basil/test"
)

var _ = Describe("GenerateInterpreter", func() {

	var registry = block.Registry{
		"block_simple":               fixtures.BlockSimpleInterpreter{},
		"block_value_required":       fixtures.BlockValueRequiredInterpreter{},
		"block_with_block":           fixtures.BlockWithBlockInterpreter{},
		"block_with_block_interface": fixtures.BlockWithBlockInterfaceInterpreter{},
		"block_with_node":            fixtures.BlockWithNodeInterpreter{},
	}

	Context("fixtures/block_simple.go", func() {
		It("should parse the input", func() {
			test.ExpectBlockToEvaluate(parser.Block(), registry)(
				`block_simple foo`,
				&fixtures.BlockSimple{IDField: "foo"},
				func(b1i interface{}, b2i interface{}, input string) {
					Expect(b1i).To(Equal(b2i), "input was %s", input)
				},
			)
		})

		It("should parse the input in short format", func() {
			test.ExpectBlockToEvaluate(parser.Block(), registry)(
				`block_simple foo "bar"`,
				&fixtures.BlockSimple{IDField: "foo", Value: "bar"},
				func(b1i interface{}, b2i interface{}, input string) {
					Expect(b1i).To(Equal(b2i), "input was %s", input)
				},
			)
		})

		It("should not parse fields with nil values", func() {
			test.ExpectBlockToEvaluate(parser.Block(), registry)(
				`block_simple foo {
					value = nil
				}`,
				&fixtures.BlockSimple{IDField: "foo"},
				func(b1i interface{}, b2i interface{}, input string) {
					Expect(b1i).To(Equal(b2i), "input was %s", input)
				},
			)
		})
	})

	Context("fixtures/block_value_required.go", func() {
		It("should parse the input in short format", func() {
			test.ExpectBlockToEvaluate(parser.Block(), registry)(
				`block_value_required foo "bar"`,
				&fixtures.BlockValueRequired{IDField: "foo", Value: "bar"},
				func(b1i interface{}, b2i interface{}, input string) {
					Expect(b1i).To(Equal(b2i), "input was %s", input)
				},
			)
		})

		It("should parse the input in short f", func() {
			test.ExpectBlockToEvaluate(parser.Block(), registry)(
				`block_value_required foo {
					value = "bar"
				}`,
				&fixtures.BlockValueRequired{IDField: "foo", Value: "bar"},
				func(b1i interface{}, b2i interface{}, input string) {
					Expect(b1i).To(Equal(b2i), "input was %s", input)
				},
			)
		})
	})

	Context("fixtures/block_with_block.go", func() {
		It("should parse the input", func() {
			test.ExpectBlockToEvaluate(parser.Block(), registry)(
				`block_with_block foo {
					block_simple bar
				}`,
				&fixtures.BlockWithBlock{IDField: "foo", Blocks: []*fixtures.BlockSimple{
					{IDField: "bar"},
				}},
				func(b1i interface{}, b2i interface{}, input string) {
					b1 := b1i.(*fixtures.BlockWithBlock)
					b2 := b2i.(*fixtures.BlockWithBlock)
					Expect(b1.IDField).To(Equal(b2.IDField), "IDField does not match, input was %s", input)
					Expect(b1.Blocks).To(Equal(b2.Blocks), "Blocks does not match, input was %s", input)
				},
			)
		})
	})

	Context("fixtures/block_with_block_interface.go", func() {
		It("should parse the input", func() {
			test.ExpectBlockToEvaluate(parser.Block(), registry)(
				`block_with_block_interface foo {
					block_simple bar
				}`,
				&fixtures.BlockWithBlockInterface{IDField: "foo", Blocks: []fixtures.BlockInterface{
					&fixtures.BlockSimple{IDField: "bar"},
				}},
				func(b1i interface{}, b2i interface{}, input string) {
					b1 := b1i.(*fixtures.BlockWithBlockInterface)
					b2 := b2i.(*fixtures.BlockWithBlockInterface)
					Expect(b1.IDField).To(Equal(b2.IDField), "IDField does not match, input was %s", input)
					Expect(b1.Blocks).To(Equal(b2.Blocks), "Blocks does not match, input was %s", input)
				},
			)
		})
	})

	Context("fixtures/block_with_node.go", func() {
		It("should parse the input", func() {
			test.ExpectBlockToEvaluate(parser.Block(), registry)(
				`block_with_node foo {
					block_simple bar
				}`,
				&fixtures.BlockWithNode{IDField: "foo"},
				func(b1i interface{}, b2i interface{}, input string) {
					b1 := b1i.(*fixtures.BlockWithNode)
					b2 := b2i.(*fixtures.BlockWithNode)
					Expect(b1.IDField).To(Equal(b2.IDField), "IDField does not match, input was %s", input)
					test.ExpectBlockNodeToEvaluate(parser.Block(), registry, b1, b1.BlockNodes[0])(
						input,
						&fixtures.BlockSimple{IDField: "bar"},
						func(b1i interface{}, b2i interface{}, input string) {
							b1 := b1i.(*fixtures.BlockSimple)
							b2 := b2i.(*fixtures.BlockSimple)
							Expect(b1.IDField).To(Equal(b2.IDField), "IDField does not match, input was %s", input)
						},
					)
				},
			)
		})
	})
})
