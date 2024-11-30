package longestsubstringwithoutrepeatingcharacters

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	highest := 0
	current := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if strings.ContainsRune(s[start:i], rune(s[i])) {
			i = start + 1
			current = 0
			start = i
		}

		current += 1

		if highest < current {
			highest = current
		}
	}

	return highest
}
