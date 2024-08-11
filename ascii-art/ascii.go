package ascii

import (
	"fmt"
	"os"
	"strings"
	fs "output/ascii-art/res"
)

func Ascii(banner string, input string) string {
	var s string
	inputLines := strings.Split(input, "\\n")
	if !fs.IsPrintable(inputLines) {
		os.Exit(1)
	}
	str := fs.Reader(banner)
	fs.Splitter(banner, str)

	// checking if the input has only newlines
	inputLines = fs.OnlyNewLine(inputLines)
	for _, value := range inputLines {
		if value == "" {
			if fs.IsOtput {
				s += "\n"
			} else {
				fmt.Println()
			}

		} else {
			if fs.IsOtput {
				s += fs.Printer(value, fs.Slice)
			} else {
				fs.Printer(value, fs.Slice)
			}

		}
	}
	fs.Slice = [][]string{}
	return s
}
