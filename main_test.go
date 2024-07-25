package main_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"ascii/asciiart"
)

func TestPrintASCIIArt(t *testing.T) {
	tests := []struct {
		name       string
		bannerFile string
		arguments  string
		want       string
		wantErr    bool
	}{
		{
			name:       "ValidInput",
			bannerFile: "standard.txt",
			arguments:  "hello",
			want: `
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
			wantErr: false,
		},
		{
			name:       "EmptyArgument",
			bannerFile: "standard.txt",
			arguments:  "",
			want:       "",
			wantErr:    false,
		},
		{
			name:       "NonASCIICharacter",
			bannerFile: "standard.txt",
			arguments:  "hellö",
			want:       "Error: Non-ASCII/printable characters found",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for output
			tmpFile, err := os.CreateTemp("", "asciiart_test")
			if err != nil {
				t.Fatalf("failed to create temporary file: %v", err)
			}
			defer os.Remove(tmpFile.Name())

			// Set the os.Args to simulate command-line arguments
			os.Args = []string{"program", fmt.Sprintf("--output=%s", tmpFile.Name())}

			// Read the banner file content
			fileContent, _ := asciiart.ReadBannerFile(tt.bannerFile)
			lines := asciiart.SplitLines(fileContent, tt.bannerFile)

			// Capture the output
			if tt.wantErr {
				// Redefine os.Exit to prevent actual exit during testing
				defer func() { recover() }()
				asciiart.WriteASCIIArt(lines, tt.arguments)
			} else {
				asciiart.WriteASCIIArt(lines, tt.arguments)
			}

			// Read content from the temporary file
			content, err := os.ReadFile(tmpFile.Name())
			if err != nil {
				t.Fatalf("failed to read content from temporary file: %v", err)
			}

			// Compare the content with the expected output
			if tt.wantErr {
				if !strings.Contains(fmt.Sprint(content), tt.want) {
					t.Errorf("WriteASCIIArt() got = %v, want %v", string(content), tt.want)
				}
			} else {
				if !strings.Contains(string(content), strings.TrimSpace(tt.want)) {
					t.Errorf("WriteASCIIArt() got = %v, want %v", string(content), tt.want)
				}
			}
		})
	}
}

func TestGetBannerFileFromArgs(t *testing.T) {
	args := []string{"Program", "--output=banner.txt", "arg2", "standard"}
	expected := "standard.txt"
	fileName := asciiart.GetBannerFileFromArgs(args)
	if fileName != expected {
		t.Errorf("Expected filename: %s, but got: %s", expected, fileName)
	}
}

func TestReadBannerFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "ValidFileStandard",
			args:    args{filename: "standard.txt"},
			wantErr: false,
		},
		{
			name:    "ValidFileShadow",
			args:    args{filename: "shadow.txt"},
			wantErr: false,
		},
		{
			name:    "ValidFileThinkertoy",
			args:    args{filename: "thinkertoy.txt"},
			wantErr: false,
		},
		{
			name:    "EmptyFilename",
			args:    args{filename: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := asciiart.ReadBannerFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBannerFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Only check content for valid files
			if !tt.wantErr {
				expectedContent, err := os.ReadFile(tt.args.filename)
				if err != nil {
					t.Fatalf("failed to read the expected content file: %v", err)
				}
				if got != string(expectedContent) {
					t.Errorf("ReadBannerFile() = %v, want %v", got, string(expectedContent))
				}
			}
		})
	}
}

func TestEscapeSequence(t *testing.T) {
	tests := []struct {
		name      string
		arguments string
		want      string
	}{
		{
			name:      "Backspace",
			arguments: "hello\\bworld",
			want:      "hellworld",
		},
		{
			name:      "Tab",
			arguments: "hello\\tworld",
			want:      "hello    world",
		},
		{
			name:      "Newline",
			arguments: "hello\\nworld",
			want:      "hello\nworld",
		},
		{
			name:      "NoEscape",
			arguments: "helloworld",
			want:      "helloworld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := asciiart.EscapeSequence(tt.arguments)
			if got != tt.want {
				t.Errorf("EscapeSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitLines(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		args       string
		wantLength int
	}{
		{
			name:       "StandardSplit",
			content:    "hello\nworld",
			args:       "standard.txt",
			wantLength: 2,
		},
		{
			name:       "ThinkertoySplit",
			content:    "hello\r\nworld",
			args:       "thinkertoy.txt",
			wantLength: 2,
		},
		{
			name:       "EmptyContent",
			content:    "",
			args:       "standard.txt",
			wantLength: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := asciiart.SplitLines(tt.content, tt.args)
			if len(got) != tt.wantLength {
				t.Errorf("SplitLines() = %v, want %v", len(got), tt.wantLength)
			}
		})
	}
}
