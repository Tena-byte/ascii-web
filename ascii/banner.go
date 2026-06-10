package ascii

import (
	"os"
	"strings"
)

func LoadBanner(filename string) map[rune][]string {

	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	}

	result := map[rune][]string{}

	data, _ := os.ReadFile("banners/" + filename)
	lines := strings.Split(string(data), "\n")

	for ch := ' '; ch <= '~'; ch++ {

		startIndex := int(ch-32)*9 + 1

		result[ch] = lines[startIndex : startIndex+8]
	}
	return result
}
