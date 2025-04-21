package analyzer

import (
	"regexp"
	"strings"
)

func AnalyzeText(text string) []int {
	sentenceEnd := regexp.MustCompile(`[.!?]\s+`)
	sentences := sentenceEnd.Split(text, -1)

	counts := []int{}
	for i, sentence := range sentences {
		if i >= 100 {
			break
		}
		words := strings.Fields(sentence)
		counts = append(counts, len(words))
	}
	return counts
}
