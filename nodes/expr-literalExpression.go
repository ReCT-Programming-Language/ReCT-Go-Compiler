package nodes

import (
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

// basic global statement member
type LiteralExpressionNode struct {
	ExpressionNode

	LiteralToken lexer.Token
	LiteralValue interface{}
	IsNative     bool
}

// implement node type from interface
func (LiteralExpressionNode) NodeType() NodeType { return LiteralExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node LiteralExpressionNode) Span() print.TextSpan {
	return node.LiteralToken.Span
}

// node print function
func (node LiteralExpressionNode) Print(indent string) {
	print.PrintC(print.Yellow, indent+"└ LiteralExpressionNode")
	fmt.Printf("%s  └ Value: %s\n", indent, node.LiteralToken.Value)
}

// "constructor" / ooga booga OOP cave man brain
func CreateLiteralExpressionNode(tok lexer.Token) LiteralExpressionNode {
	return LiteralExpressionNode{
		LiteralToken: tok,
		LiteralValue: tok.RealValue,
		IsNative:     false,
	}
}

func CreateNativeLiteralExpressionNode(tok lexer.Token) LiteralExpressionNode {
	return LiteralExpressionNode{
		LiteralToken: tok,
		LiteralValue: tok.RealValue,
		IsNative:     true,
	}
}
