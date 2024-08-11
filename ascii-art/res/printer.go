package fs

import "fmt"

var IsOtput bool

func Printer(inputLine string, slice [][]string) string {
	var s string
	for j := 0; j < 8; j++ {
		for _, char := range inputLine {
			index := int(char) - 32
			if IsOtput {
				s += slice[index][j]
			} else {
				fmt.Print(slice[index][j])
			}

		}
		if IsOtput {
			s += "\n"
		} else {
			fmt.Println()
		}

	}
	return s
}
