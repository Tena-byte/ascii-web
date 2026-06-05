package ascii

import "strings"

func Render(text string, font map[rune][]string) string {

	result := make([]string, 8)

	for _, ch := range text {

		for row := 0; row < 8; row++ {

			result[row] += font[ch][row]
		}
	}
	return strings.Join(result, "\n")
}
