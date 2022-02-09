package main

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/evaluator"
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/parser"
	"ReCT-Go-Compiler/print"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

/* cli.go handles flags and command line arguments for the project
 * Everything is documented below, this was moved away from its own package
 * formally cli/cli.go because it makes more sense for the flags to be processed
 * inside the main package - you gain no benefit from having cli as its own package.
 */

// Defaults
// These are the default values of all flags
// The values are set using the flag library, but they are also commented below
var helpFlag bool      // false -h
var interpretFlag bool // true  -i
var showVersion bool   // false -v
var fileLog bool       // false -l
var debug bool         // -xx
var tests bool         // Just for running test file like test.rct ( -t )
var files []string

// Constants that are used throughout code
// Should be updated when necessary
const executableName string = "rgoc"                         // in case we change it later
const discordInvite string = "https://discord.gg/kk9MsnABdF" // infinite link
const currentVersion string = "1.1"

// Init initializes and processes (parses) compiler flags
func Init() {
	flag.BoolVar(&helpFlag, "h", false, "Shows this help message")
	flag.BoolVar(&interpretFlag, "i", true, "Enables interpreter mode, source code will be interpreted instead of compiled.")
	flag.BoolVar(&showVersion, "v", false, "Shows current ReCT version the compiler supports")
	flag.BoolVar(&fileLog, "l", false, "Logs process information in a log file")
	flag.BoolVar(&debug, "xx", false, "Shows brief process information in the command line")
	// Test (-t) will not be in the help message as it's only really going ot be used for testing compiler features.
	flag.BoolVar(&tests, "t", false, "For compiler test files (developers only)")
	files = flag.Args() // Other arguments like executable name or files
	flag.Parse()
}

// ProcessFlags goes through each flag and decides how they have an effect on the output of the compiler
func ProcessFlags() {
	// Mmm test has the highest priority
	if tests {
		RunTests()
		return // returns to main
	}

	// Show version has higher priority than help menu
	if showVersion {
		Version()
		return // returns to main
	}
	// If they use "-h" or only enter the executable name "rgoc"
	// Show the help menu because they're obviously insane.
	if helpFlag || len(files) <= 1 {
		Help()
		return // returns to main
	}
	if interpretFlag {
		InterpretFile(files[1])
	}
}

// InterpretFile runs everything to interpret the files, currently only supports up to one file
func InterpretFile(file string) {
	print.WriteC(print.Green, "-> Lexing...  ")
	tokens := lexer.Lex(file)
	print.PrintC(print.Green, "Done!")

	print.WriteC(print.Yellow, "-> Parsing... ")
	members := parser.Parse(tokens)
	print.PrintC(print.Green, "Done!")

	print.WriteC(print.Red, "-> Binding... ")
	boundProgram := binder.BindProgram(members)
	print.PrintC(print.Green, "Done!")
	//boundProgram.Print()

	print.PrintC(print.Cyan, "-> Evaluating!")
	evaluator.Evaluate(boundProgram)
}

// RunTests runs all the test files in /tests/
func RunTests() {
	files, err := ioutil.ReadDir("tests")
	if err != nil {
		// better error later
		print.PrintC(print.DarkRed, "ERROR: failed reading /tests/ directory!")
	}
	tests := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".rct") {
			tests = append(tests, file.Name())
		}
	}
	for _, test := range tests {
		print.PrintC(
			print.Cyan,
			fmt.Sprintf("\nTesting test file \"%s\":", test),
		)
		// forgot to actually run the file lol
		go InterpretFile("tests/" + test)

	}
}

// Help shows help message (pretty standard nothing special)
func Help() {
	header := "ReCT Go Compiler v" + currentVersion
	lines := strings.Repeat("-", len(header))

	fmt.Println(lines)
	fmt.Println(header)
	fmt.Println(lines)

	fmt.Print("\nUsage: ")
	print.PrintC(print.Green, "rgoc <file> [options]\n")
	fmt.Println("<file> can be the path to any ReCT file (.rct)")
	fmt.Println("\n[Options]")

	helpSegments := []HelpSegment{
		{"Help", executableName + " -h", "disabled (default)", "Shows this help message!"},
		{"Interpret", executableName + " -i", "enabled (default)", "Enables interpreter mode, source code will be interpreted instead of compiled."},
		{"File logging", executableName + " -l", "disabled (default)", "Logs process information in a log file"},
		{"Debug", executableName + " -xx", "disabled (default)", "Shows brief process information in the command line"},
	}

	p0, p1, p2, p3 := findPaddings(helpSegments)

	for _, segment := range helpSegments {
		segment.Print(p0, p1, p2, p3)
	}

	fmt.Println("")
	print.PrintCF(print.Gray, "Still having troubles? Get help on the offical Discord server: %s!\n", discordInvite)
}

// Version Shows the current compiler version
func Version() {
	fmt.Println("ReCT Go Compiler")
	fmt.Print("ReCT version: ")
	print.PrintC(print.Blue, currentVersion)
	fmt.Printf("\nFor more informatin, why not join the discord? %s\n\n", discordInvite)
}

type HelpSegment struct {
	Command      string
	Example      string
	DefaultValue string
	Explaination string
}

func (seg *HelpSegment) Print(p0 int, p1 int, p2 int, p3 int) {
	print.WriteCF(print.Cyan, "%-*s", p0, seg.Command)
	print.WriteC(print.DarkGray, ":")
	print.WriteCF(print.Blue, " %-*s", p1, seg.Example)
	print.WriteC(print.DarkGray, ":")
	print.WriteCF(print.Yellow, " %-*s", p2, seg.DefaultValue)
	print.WriteC(print.DarkGray, ":")
	print.WriteCF(print.Green, " %-*s", p3, seg.Explaination)
	fmt.Println("")
}

func findPaddings(segments []HelpSegment) (int, int, int, int) {
	p0 := 0
	p1 := 0
	p2 := 0
	p3 := 0

	for _, segment := range segments {
		if len(segment.Command) > p0 {
			p0 = len(segment.Command)
		}
		if len(segment.Example) > p1 {
			p1 = len(segment.Example)
		}
		if len(segment.DefaultValue) > p2 {
			p2 = len(segment.DefaultValue)
		}
		if len(segment.Explaination) > p3 {
			p3 = len(segment.Explaination)
		}
	}

	return p0 + 1, p1 + 1, p2 + 1, p3 + 1
}
