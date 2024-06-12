package asciiart

import "strings"

func SplitLines(content string, args string) []string {
	if args == "thinkertoy.txt" {
		return strings.Split(content, "\r\n")
	}
	return strings.Split(content, "\n")
}
