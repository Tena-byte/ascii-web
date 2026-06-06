package ascii

import (
	"os"
	"strings"
)



func LaodBanner(filename string) map[rune][]string{

	Filename := strings.ToLower(filename)

	if !strings.HasSuffix(Filename, ".txt"){
		Filename += ".txt"
	}

	result := map[rune][]string{}

	data, _ := os.ReadFile("banners/" + Filename)

	lines := strings.Split(string(data), "\n")

	
	for ch := ' '; ch <= '~'; ch++ {

		startIndex := int(ch-' ')*9 + 1

		result[ch] = lines[startIndex : startIndex+8]
	}

	return result
}