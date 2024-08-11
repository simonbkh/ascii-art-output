package fs

import "fmt"

func IsPrintable(sl []string) bool {
	for _, word := range sl {
		for _, char := range word {
			if char < 32 || char > 126 {
				fmt.Printf("Character '%c' is not a printable ASCII character\n", char)
				return false
			}
		}
	}
	return true
}
