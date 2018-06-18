package ast

import (
	"github.com/nickwanninger/act/pkg/types"
)

func (p *Parser) parseNumericExpr() Node {
	t, val := p.token.InferType()

	if t == types.DefaultIntType {
		n := intNode{}
		n.NodeType = nodeInt
		n.Value = val.(int64)
		p.next()
		return n
	}

	if t == types.DefaultFloatType {
		n := floatNode{}
		n.NodeType = nodeFloat
		n.Value = val.(float64)
		p.next()
		return n
	}

	Error(p.token, "invalid number syntax")

	return nil
}