package asciiart

import (
	"fmt"
	"os"
)

// ReadBannerFile reads and returns the contents of the banner file based on the provided filename
func ReadBannerFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileSize := fileInfo.Size()
	if filename == "standard.txt" && fileSize != 6623 {
		fmt.Println("file is corrupted")
		os.Exit(0)
	}

	if filename == "shadow.txt" && fileSize != 7463 {
		fmt.Println("file is corrupted")
		os.Exit(0)
	}

	if filename == "thinkertoy.txt" && fileSize != 5558 {
		fmt.Println("file is corrupted")
		os.Exit(0)
	}
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
