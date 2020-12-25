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

	topCount := 10
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

	if len(sorted) < 10 {
		topCount = len(sorted)
	}

	var top = make([]string, topCount)
	for i, val := range sorted[:topCount] {
		top[i] = val.Word
	}

	return top
}

func sortByFreq(wordFreq map[string]int) []WordFreq {
	var sorted = make([]WordFreq, len(wordFreq))

	i := 0
	for word, freq := range wordFreq {
		sorted[i] = WordFreq{word, freq}
		i++
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Freq > sorted[j].Freq
	})

	return sorted
}
