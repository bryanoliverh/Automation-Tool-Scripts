package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Define command line arguments
	findPtr := flag.String("T", "", "Text to find")
	replacePtr := flag.String("R", "", "Text to replace with")
	filePtr := flag.String("file", "", "File to modify")
	fileAltPtr := flag.String("f", "", "Alternative file to modify")
	casePtr := flag.String("C", "ci", "Specify case sensitivity (ci for case insensitive, cs for case sensitive)")
	outputPtr := flag.String("O", "output.txt", "Output file")
	flag.StringVar(findPtr, "t", "", "Text to find (shorthand)")
	flag.StringVar(replacePtr, "r", "", "Text to replace with (shorthand)")
	flag.StringVar(fileAltPtr, "F", "", "Alternative file to modify (shorthand)")
	flag.StringVar(casePtr, "case", "ci", "Specify case sensitivity (shorthand)")
	flag.StringVar(outputPtr, "out", "output.txt", "Output file (shorthand)")

	// Define usage information
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <file>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Find and replace text in a file.")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Fprintln(os.Stderr, "  -T, --t TEXT         Text to find")
		fmt.Fprintln(os.Stderr, "  -R, --r TEXT         Text to replace with")
		fmt.Fprintln(os.Stderr, "  -file FILE           File to modify")
		fmt.Fprintln(os.Stderr, "  -f FILE              Alternative file to modify (shorthand)")
		fmt.Fprintln(os.Stderr, "  -C, --case CASE      Specify case sensitivity (ci for case insensitive, cs for case sensitive)")
		fmt.Fprintln(os.Stderr, "  -O, --out FILE       Output file")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Examples:")
		fmt.Fprintln(os.Stderr, "  findreplace -T \"foo\" -R \"bar\" input.txt")
		fmt.Fprintln(os.Stderr, "  findreplace --t \"foo\" --r \"bar\" --out output.txt --case cs --file input.txt")
	}

	flag.Parse()

	// Check if both -T and -R are provided
	if *findPtr == "" || *replacePtr == "" {
		fmt.Println("Both -T and -R are required.")
		os.Exit(1)
	}

	// Check if either -file, -f or input file path are provided
	var file string
	if *filePtr != "" {
		file = *filePtr
	} else if *fileAltPtr != "" {
		file = *fileAltPtr
	} else if len(flag.Args()) == 1 {
		file = flag.Arg(0)
	} else {
		fmt.Println("Either -file, -f or input file path are required.")
		os.Exit(1)
	}

	// Read input file
	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	// Find and replace text
	var output string
	if *casePtr == "ci" {
		output = strings.ReplaceAll(strings.ToLower(string(input)), strings.ToLower(*findPtr), *replacePtr)
	} else if *casePtr == "cs" {
		output = strings.ReplaceAll(string(input), *findPtr, *replacePtr)
	} else {
		fmt.Println("Specify case sensitivity with -C or -case (ci or cs).")
		os.Exit(1)
	}

	// Write output file
	err = ioutil.WriteFile(*outputPtr, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Find and replace successful!")
	fmt.Printf("Output written to %s\n", *outputPtr)
}
