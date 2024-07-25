package asciiart

import (
	"fmt"
	"os"
	"strings"
)

// read the string from the arguments
func EscapeSequence(arguments string) string {
	// arguments := args[2]

	// Handling backspace tabs
	arguments = strings.ReplaceAll(arguments, "\\b", "\b")

	for {
		index := strings.Index(arguments, "\b")

		if index == -1 {
			break
		}
		if index > 0 {
			arguments = arguments[:index-1] + arguments[index+1:]
		} else {
			arguments = arguments[index+1:]
		}
	}

	if strings.Contains(arguments, "\\t") {
		arguments = strings.ReplaceAll(arguments, "\\t", "    ")
	} else if strings.Contains(arguments, "\\n") {
		arguments = strings.ReplaceAll(arguments, "\\n", "\n")
	} else if strings.Contains(arguments, "\\a") || strings.Contains(arguments, "\\r") || strings.Contains(arguments, "\\'") || strings.Contains(arguments, "\\v") || strings.Contains(arguments, "\\f") || strings.Contains(arguments, "\\x20") || strings.Contains(arguments, "\\x0c") {
		fmt.Println("Error: Special Character found")
		os.Exit(0)
	}
	return strings.ReplaceAll(arguments, "\\n", "\n")
}
