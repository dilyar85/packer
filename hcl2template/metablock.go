package hcl2template

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/packer/hcl2template/addrs"
	"github.com/zclconf/go-cty/cty"
)

// MetaBlock defines funcs that are common to all HCL2 executable blocks.
type MetaBlock interface {
	// References returns the *addrs.Reference of everything being referenced in
	// this block.
	References() ([]*addrs.Reference, hcl.Diagnostics)

	// Evaluate and validate the expressions for this block and properly
	// configure it.
	Evaluate(ctx *hcl.EvalContext) hcl.Diagnostics

	// Value of the block. It is probably unknown or not fully known before Run
	// is called.
	Value() cty.Value

	// Expected type of Value, when unknown: cty.DynamicPseudoType is returned.
	Type() cty.Type
}

var (
	_ MetaBlock = &Variable{}
)
