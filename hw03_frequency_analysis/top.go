package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

type WordFreq struct {
	Word string
	Freq int
}

func Top10(text string) []string {
	if text == "" {
		return []string{}
	}

	words := strings.Fields(text)
	sort.Strings(words)

	wordFreqMap := make(map[string]int)
	for _, val := range words {
		cnt, exists := wordFreqMap[val]
		if exists {
			wordFreqMap[val] = cnt + 1
		} else {
			wordFreqMap[val] = 1
		}
	}

	sorted := sortByFreq(wordFreqMap)

	var top = make([]string, 10)
	for i, val := range sorted[:10] {
		top[i] = val.Word
	}

	return top
}

func sortByFreq(wordFreq map[string]int) []WordFreq {
	var sorted = make([]WordFreq, len(wordFreq))

	for word, freq := range wordFreq {
		sorted = append(sorted, WordFreq{word, freq})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Freq > sorted[j].Freq
	})

	return sorted
}
