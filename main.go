package main

import (
	"fmt"
	"os"
	"output/ascii-art"
	fs "output/ascii-art/res"
	"strings"
)

var (
	outputFile string
)

func formatError() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}

func optionFlag(k string) bool {

	if strings.HasPrefix(string(k), "--output=") {
		outputFile = strings.TrimPrefix(string(k), "--output=")
	} else {
		formatError()
		return false
	}
return true
}

func main() {
	var banner string
	var input string
	if len(os.Args) != 1 && strings.HasPrefix(os.Args[1], "--output") {
		if len(os.Args) < 3 || len(os.Args) > 4 {
			formatError()
			return
		} else {
			if len(os.Args) == 3 {
				input = os.Args[2]
				banner = "standard"
				fs.IsOtput = true
				if !optionFlag(os.Args[1]){
					return
				}
				
			} else {
				input = os.Args[2]
				banner = os.Args[3]
				if !optionFlag(os.Args[1]){
					return
				}
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
	fullOutput := ascii.Ascii(banner, input)
	if fs.IsOtput {
		if !strings.HasSuffix(outputFile, ".txt") {
			fmt.Println("Error: File must be .txt extension")
			return
		}
	
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
