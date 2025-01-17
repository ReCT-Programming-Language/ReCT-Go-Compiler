package nodes

import print2 "github.com/ReCT-Lang/ReCT-Go-Compiler/print"

// very cool interface for creating syntax nodes
type SyntaxNode interface {
	NodeType() NodeType
	Span() print2.TextSpan // exact text position of this node
	Print(indent string)
	// only type atm, might contain more stuff like text-location later
}

// very cool interface for creating statements
// (ik this isnt very revolutionary but i just like to organise stuff)
type StatementNode interface {
	SyntaxNode
}

// very cool interface for creating members
// (again, organisation.)
type MemberNode interface {
	SyntaxNode
}

// very cool interface for creating expressions
// (craaaaazy ikr)
type ExpressionNode interface {
	SyntaxNode
}

// cool node type Enum straight up stolen from ReCT v1.0
type NodeType string

const (
	// i am basing these objects off of the rect 1.0 source
	// -> https://github.com/RedCubeDev-ByteSpace/ReCT/tree/834776cbf0ad97da0e6441835f1bc19d903f115b/ReCT/CodeAnalysis/Syntax

	// Members
	// -------
	GlobalStatement             NodeType = "Global Statement"
	FunctionDeclaration         NodeType = "Function Declaration"
	ExternalFunctionDeclaration NodeType = "External Function Declaration"
	ClassDeclaration            NodeType = "Class Declaration"
	StructDeclaration           NodeType = "Struct Declaration"
	EnumDeclaration             NodeType = "Enum Declaration"
	PackageReference            NodeType = "Package Reference"
	PackageAlias                NodeType = "Package Alias"
	PackageUse                  NodeType = "Package Use"

	// General
	// -------
	Parameter  NodeType = "Parameter"
	TypeClause NodeType = "Type Clause"

	// Statements
	// ----------
	BlockStatement      NodeType = "Block Statement"
	VariableDeclaration NodeType = "Variable Declaration"
	IfStatement         NodeType = "If Statement"
	ElseClause          NodeType = "Else Clause"
	ReturnStatement     NodeType = "Return Statement"
	ForStatement        NodeType = "For Statement"
	WhileStatement      NodeType = "While Statement"
	BreakStatement      NodeType = "Break Statement"
	ContinueStatement   NodeType = "Continue Statement"
	FromToStatement     NodeType = "FromTo Statement"
	ExpressionStatement NodeType = "Expression Statement"

	// Expressions
	// -----------
	LiteralExpression              NodeType = "Literal Expression"
	ParenthesisedExpression        NodeType = "Parenthesised Expression"
	NameExpression                 NodeType = "Name Expression"
	AssignmentExpression           NodeType = "Assignment Expression"
	CallExpression                 NodeType = "Call Expression"
	PackageCallExpression          NodeType = "PackageCall Expression"
	UnaryExpression                NodeType = "Unary Expression"
	BinaryExpression               NodeType = "Binary Expression"
	VariableEditorExpression       NodeType = "VariableEditor Expression"
	TypeCallExpression             NodeType = "TypeCall Expression"
	ClassFieldAccessExpression     NodeType = "ClassFieldAccess Expression"
	ClassFieldAssignmentExpression NodeType = "ClassFieldAssignment Expression"
	ArrayAccessExpression          NodeType = "ArrayAccess Expression"
	ArrayAssignmentExpression      NodeType = "ArrayAssignment Expression"
	MakeExpression                 NodeType = "Make Expression"
	MakeArrayExpression            NodeType = "MakeArray Expression"
	TernaryExpression              NodeType = "Ternary Expression"
	ReferenceExpression            NodeType = "Reference Expression"
	DereferenceExpression          NodeType = "Dereference Expression"
	MakeStructExpression           NodeType = "MakeStruct Expression"
	LambdaExpression               NodeType = "Lambda Expression"
	ThisExpression                 NodeType = "This Expression"
)
