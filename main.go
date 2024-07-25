package main

import (
	"fmt"
	"os"

	"ascii/asciiart"
)

func main() {
	// Get command-line arguments
	args := os.Args

	// Checking if there are enough arguments in the command line
	if len(args) < 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}

	// Get banner file from command line arguments
	bannerFile := asciiart.GetBannerFileFromArgs(args)

	// Read the content of the selected banner file
	fileContent, err := asciiart.ReadBannerFile(bannerFile)
	if err != nil {
		fmt.Println("Error: Missing file")
		os.Exit(0)
	}

	// Split the content into lines
	lines := asciiart.SplitLines(fileContent, bannerFile)
	if len(lines) != 856 {
		fmt.Println("corrupt file")
		return
	}

	input := args[2]
	// Get the parsed argument
	argument := asciiart.EscapeSequence(input)

	// Print the ASCII art
	asciiart.WriteASCIIArt(lines, argument)
}
