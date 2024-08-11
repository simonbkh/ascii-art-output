package main

import (
	"fmt"
	"os"
	"strings"
	"output/ascii-art"
	fs "output/ascii-art/res"
)

var (
	outputFile string
)

func formatError() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}

func optionFlag() {

	if strings.HasPrefix(string(os.Args[1]), "--output=") {
		outputFile = strings.TrimPrefix(string(os.Args[1]), "--output=")
	} else {
		formatError()
	}

}

func main() {
	var banner string
	var input string
	if strings.HasPrefix(os.Args[1], "--output=") {
		if len(os.Args) < 3 || len(os.Args) > 4 {
			formatError()
			return
		} else {
			if len(os.Args) == 3 {
				input = os.Args[2]
				banner = "standard"
				fs.IsOtput = true
				optionFlag()
			} else {
				input = os.Args[2]
				banner = os.Args[3]
				optionFlag()
				fs.IsOtput = true
			}
		}
	} else {
		if len(os.Args) == 2 {
			input = os.Args[1]
			banner = "standard"
		} else if len(os.Args) == 3 {
			input = os.Args[1]
			banner = os.Args[2]
		} else {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			return
		}
	}
	if input == "" {
		return
	}

	if fs.IsOtput {
		if !strings.HasSuffix(outputFile, ".txt") {
			fmt.Println("Error: File must be .txt extension")
			return
		}
		fullOutput := ascii.Ascii(banner, input)
		err := os.WriteFile(outputFile, []byte(fullOutput), 0o644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	// v := ascii.Ascii(banner, input)
	// if v != "" {
	// 	os.WriteFile(outputFile, []byte(v), 0o644)
	// }

}
