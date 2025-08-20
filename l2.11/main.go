package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	firstEncounteredWords := make(map[string]string)
	res := make(map[string][]string)

	for _, word := range words {
		word = strings.ToLower(word)
		runes := []rune(word)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedWord := string(runes)

		if _, ok := firstEncounteredWords[sortedWord]; !ok {
			firstEncounteredWords[sortedWord] = word
		}

		key := firstEncounteredWords[sortedWord]
		fmt.Println(key)
		if anagrams, exists := res[key]; exists {
			anagrams = append(anagrams, word)
			res[key] = anagrams
		} else {
			res[key] = []string{word}
		}
	}
	return res
}

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	res := findAnagrams(input)

	for word, anagrams := range res {
		if len(anagrams) > 1 {
			fmt.Println(word, anagrams)
		}
	}
}
