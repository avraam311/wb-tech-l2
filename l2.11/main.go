package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	firstEncounteredWords := make(map[string]string)
	groups := make(map[string][]string)

	for _, word := range words {
		lower := strings.ToLower(word)
		runes := []rune(lower)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedWord := string(runes)

		if _, ok := firstEncounteredWords[sortedWord]; !ok {
			firstEncounteredWords[sortedWord] = lower
		}

		key := firstEncounteredWords[sortedWord]
		groups[key] = append(groups[key], lower)
	}

	result := make(map[string][]string)
	for key, anagrams := range groups {
		if len(anagrams) > 1 {
			sort.Strings(anagrams)
			result[key] = anagrams
		}
	}

	return result
}

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	res := findAnagrams(input)

	for word, anagrams := range res {
		fmt.Println(word, anagrams)
	}
}
