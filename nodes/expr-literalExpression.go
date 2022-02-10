package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
	"fmt"
)

// basic global statement member
type LiteralExpressionNode struct {
	ExpressionNode

	LiteralToken lexer.Token
	LiteralValue interface{}
}

// implement node type from interface
func (LiteralExpressionNode) NodeType() NodeType { return LiteralExpression }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node LiteralExpressionNode) Position() (int, int, int) {
	length := node.LiteralToken.Column + len(node.LiteralValue.(string))
	return node.LiteralToken.Line, node.LiteralToken.Column, length
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
	}
}
