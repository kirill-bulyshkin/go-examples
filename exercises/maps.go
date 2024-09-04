package exercises

import (
	"strings"

	"golang.org/x/tour/wc"
)

// Exercise: Maps

// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.
// You might find strings.Fields helpful.

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordsMap := make(map[string]int)

	for _, word := range words {
		_, ok := wordsMap[word]
		if !ok {
			wordsMap[word] = 1
		} else {
			wordsMap[word]++
		}
	}
	return wordsMap
}

func WordCountTest() {
	wc.Test(WordCount)
}
