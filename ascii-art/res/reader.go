package fs

import (
	"fmt"
	"os"
	"os/exec"
)

func Reader(banner string) string {
	// if reading the bannner fails or the file got altered in a way that its gonna ruin the ascii-art it will be  fetched and  that will secure normal running of the program
	content, err := os.ReadFile(banner + ".txt")

	if err != nil || (banner == "standard" && len(content) != 6623) || (banner == "shadow" && len(content) != 7463) || (banner == "thinkertoy" && len(content) != 5558) {
		fmt.Println(len(content))
		url := "https://raw.githubusercontent.com/01-edu/public/master/subjects/ascii-art/" + banner + ".txt"
		cmd := exec.Command("curl", "-o", banner+".txt", url)
		_, er := cmd.Output()
		if er != nil {
			fmt.Println("Error fetching url:", er)
			os.Exit(1)
		}
		content, _ = os.ReadFile(banner + ".txt")
	}
	return string(content)
}
