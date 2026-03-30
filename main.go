package main

import (
	"flag"
	"fmt"
)

func main() {
	length := flag.Int("length", 16, "Length of password")
	count := flag.Int("count", 1, "Number of passwords")
	noSymbols := flag.Bool("no-symbols", false, "Exclude symbols")
	requireEach := flag.Bool("require-each-class", false, "Require each character class")
	showEntropy := flag.Bool("show-entropy", false, "Show entropy")

	flag.Parse()

	fmt.Println(*length, *count, *noSymbols, *requireEach, *showEntropy)
}
