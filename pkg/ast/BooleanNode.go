package ast

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// BooleanNode is an integer literal
type BooleanNode struct {
	NodeType
	TokenReference
	Accessable

	Value string
}

// NameString implements Node.NameString
func (n BooleanNode) NameString() string { return "BooleanNode" }

// Codegen implements Node.Codegen for BooleanNode
func (n BooleanNode) Codegen(prog *Program) (value.Value, error) {
	options := map[string]bool{
		"true":  true,
		"false": false,
	}
	return constant.NewBool(options[n.Value]), nil
}

func (n BooleanNode) String() string {
	return n.Value
}

// GenAccess implements Accessable.GenAccess
func (n BooleanNode) GenAccess(prog *Program) (value.Value, error) {
	return n.Codegen(prog)
}
