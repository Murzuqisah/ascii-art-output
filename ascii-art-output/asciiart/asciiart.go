package asciiart

import (
	"fmt"
	"os"
	"strings"
)

// PrintASCIIArt prints the string based on the ASCII character range to a file
func WriteASCIIArt(lines []string, arguments string) {
	argument := strings.Split(arguments, "\n")
	count := 0
	if len(os.Args[1]) < 9 {
		fmt.Println("Usage: go run . --output=banner.txt <string> <bannerfile>\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	outputFile := os.Args[1]
	Prefix := "--output="
	if !strings.HasPrefix(outputFile, Prefix) {
		fmt.Println("Error: The argument does not have the required prefix")
		os.Exit(0)
	} else {
		outputFile := os.Args[1][9:]
		Suffix := ".txt"
		if !strings.HasSuffix(outputFile, Suffix) {
			fmt.Println("Error: The argument does not have the required suffix")
			os.Exit(0)
		} else {
			var outputContent strings.Builder
			for _, arg := range argument {
				for _, chr := range arg {
					if chr < 32 || chr > 126 {
						fmt.Println("Error: Non-ASCII/printable characters found")
						os.Exit(0)
					}
				}
				if arg == "" {
					count++
					if count < len(argument) {
						outputContent.WriteString("\n")
					}
				} else {
					for i := 0; i < 8; i++ {
						for _, value := range arg {
							start := int(value-32)*9 + 1
							outputContent.WriteString(lines[start+i])
						}
						outputContent.WriteString("\n")
					}
				}
			}
			err := os.WriteFile(outputFile, []byte(outputContent.String()), 0o644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				os.Exit(1)
			}
		}
	}
}
