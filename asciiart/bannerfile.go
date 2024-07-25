package asciiart

import (
	"fmt"
	"os"
)

// GetBannerFileFromArgs determines the filename from the command line arguments and returns it
func GetBannerFileFromArgs(args []string) string {
	if len(args) == 4 {
		if args[3] == "shadow" || args[3] == "shadow.txt" {
			return "shadow.txt"
		} else if args[3] == "thinkertoy" || args[3] == "thinkertoy.txt"{
			return "thinkertoy.txt"
		} else if args[3] == "standard" || args[3] == "standard.txt"{
			return "standard.txt"
		} else {
			fmt.Println(`Unidentified file, Usage : go run . --output=banner.txt "text" bannerfile`)
			os.Exit(0)
		}
	}
	if len(args) < 3 {
		fmt.Println("Usage: go run . --output=banner.txt [STRING] [BANNER]")
		os.Exit(0)
	}
	return "standard.txt"
}
