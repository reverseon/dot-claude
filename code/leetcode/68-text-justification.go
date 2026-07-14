package main

import (
	"strings"
)

func justifyLine(words []string, maxWidth int, isLastLine bool) string {
	wordLen := 0
	for _, w := range words {
		wordLen += len(w)
	}

	totalSpaces := maxWidth - wordLen
	gaps := len(words) - 1

	if isLastLine || gaps == 0 {
		// Left-align: single space between words, pad right
		line := strings.Join(words, " ")
		return line + strings.Repeat(" ", maxWidth-len(line))
	}

	// Fully justify: distribute spaces evenly, extra spaces go left-to-right
	spacesPerGap := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	var result strings.Builder
	for i, word := range words {
		result.WriteString(word)
		if i < gaps {
			spaces := spacesPerGap
			if i < extraSpaces {
				spaces++
			}
			result.WriteString(strings.Repeat(" ", spaces))
		}
	}

	return result.String()
}

func fullJustify(words []string, maxWidth int) []string {
	// Phase 1: Group words into lines
	var lines [][]string
	var line []string
	var curLen int

	for _, word := range words {
		if curLen+len(word)+len(line) > maxWidth {
			lines = append(lines, line)
			line, curLen = []string{}, 0
		}
		line = append(line, word)
		curLen += len(word)
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}

	// Phase 2: Justify each line
	var result []string
	for i, line := range lines {
		result = append(result, justifyLine(line, maxWidth, i == len(lines)-1))
	}
	return result
}
