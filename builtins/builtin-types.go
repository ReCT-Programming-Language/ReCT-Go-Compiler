package builtins

import "ReCT-Go-Compiler/symbols"

var (
	Void   = symbols.CreateTypeSymbol("void", make([]symbols.TypeSymbol, 0), false, false)
	Bool   = symbols.CreateTypeSymbol("bool", make([]symbols.TypeSymbol, 0), false, false)
	Byte   = symbols.CreateTypeSymbol("byte", make([]symbols.TypeSymbol, 0), false, false)
	Int    = symbols.CreateTypeSymbol("int", make([]symbols.TypeSymbol, 0), false, false)
	UInt   = symbols.CreateTypeSymbol("uint", make([]symbols.TypeSymbol, 0), false, false)
	Long   = symbols.CreateTypeSymbol("long", make([]symbols.TypeSymbol, 0), false, false)
	ULong  = symbols.CreateTypeSymbol("ulong", make([]symbols.TypeSymbol, 0), false, false)
	Float  = symbols.CreateTypeSymbol("float", make([]symbols.TypeSymbol, 0), false, false)
	Double = symbols.CreateTypeSymbol("double", make([]symbols.TypeSymbol, 0), false, false)
	String = symbols.CreateTypeSymbol("string", make([]symbols.TypeSymbol, 0), true, false)
	Any    = symbols.CreateTypeSymbol("any", make([]symbols.TypeSymbol, 0), true, false)

	// lambda/functionExpression/action/etc
	Action = symbols.CreateTypeSymbol("action", make([]symbols.TypeSymbol, 0), false, false)

	// threads
	Thread = symbols.CreateTypeSymbol("thread", make([]symbols.TypeSymbol, 0), true, false)

	// generic array types so the emitter has something to work with
	Array  = symbols.CreateTypeSymbol("array", make([]symbols.TypeSymbol, 0), true, false)
	PArray = symbols.CreateTypeSymbol("parray", make([]symbols.TypeSymbol, 0), true, false)

	// spoopy
	Pointer = symbols.CreateTypeSymbol("pointer", make([]symbols.TypeSymbol, 0), false, false)

	// the cursed ones
	Error    = symbols.CreateTypeSymbol("?", make([]symbols.TypeSymbol, 0), false, false)
	Identity = symbols.CreateTypeSymbol("¯\\_(ツ)_/¯", make([]symbols.TypeSymbol, 0), false, false)

	Types = []symbols.TypeSymbol{
		Void, Bool, Byte, Int, Long, UInt, ULong, Float, Double, String, Any, Action, Array, PArray, Pointer, Thread,
	}
)
