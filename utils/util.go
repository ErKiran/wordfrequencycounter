package utils

import (
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func sentancesToArrayOfWords(word string) ([]string, error) {
	re, err := regexp.Compile("[^a-zA-Z0-9']+")
	if err != nil {
		return nil, err
	}
	word = re.ReplaceAllString(word, " ")

	words := strings.Split(strings.ToLower(word), " ")

	return words, err
}

func TextToWordCount(text string) ([]WordCount, error) {
	words, err := sentancesToArrayOfWords(text)
	if err != nil {
		return nil, err
	}
	wordMap := make(map[string]int)

	for _, v := range words {
		if len(v) > 0 {
			wordMap[v]++
		}
	}

	var wordCounts []WordCount
	for k, v := range wordMap {
		wordCounts = append(wordCounts, WordCount{
			Word:  k,
			Count: v,
		})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})
	return wordCounts, nil
}
