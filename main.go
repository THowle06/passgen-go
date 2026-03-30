package main

import (
	"flag"
	"fmt"
	"passgen-go/generator"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	length := flag.Int("length", 16, "Length of password")
	count := flag.Int("count", 1, "Number of passwords")
	noSymbols := flag.Bool("no-symbols", false, "Exclude symbols")
	requireEach := flag.Bool("require-each-class", false, "Require each character class")
	showEntropy := flag.Bool("show-entropy", false, "Show entropy")
	clipboardFlag := flag.Bool("clipboard", false, "Copy to clipboard")

	flag.Parse()

	if *count <= 0 {
		panic("count must be > 0")
	}

	enabled := 3
	if !*noSymbols {
		enabled++
	}

	if *requireEach && *length < enabled {
		panic("length too short for required character classes")
	}

	policy := generator.Policy{
		Length:           *length,
		IncludeUpper:     true,
		IncludeLower:     true,
		IncludeDigits:    true,
		IncludeSymbols:   !*noSymbols,
		RequireEachClass: *requireEach,
	}

	var results []string

	for i := 0; i < *count; i++ {
		pwd := generator.Generate(policy)
		results = append(results, pwd)
	}

	for _, pwd := range results {
		fmt.Println(pwd)
	}

	if *showEntropy {
		entropy := generator.CalculateEntropy(policy)
		strength := generator.ClassifyEntropy(entropy)
		fmt.Printf("\nEntropy: %.2f bits (%s)\n", entropy, strength)
	}

	if *clipboardFlag {
		output := strings.Join(results, "\n")

		err := clipboard.WriteAll(output)
		if err != nil {
			fmt.Println("Clipboard error:", err)
		}
	}
}
