package main

import (
	"fmt"
	"regexp"
	"strings"
	"wordcount/api"
)

func main() {

	api.Run()

	word := "Love isn't always a ray of sunshine. That's what the older girls kept telling her when she said she had found the perfect man. She had thought this was simply bitter talk on their part since they had been unable to find true love like hers. But now she had to face the fact that they may have been right. Love may not always be a ray of sunshine. That is unless they were referring to how the sun can burn."

	words, err := helper(word)
	if err != nil {
		fmt.Println(err)
	}
	wordMap := make(map[string]int)

	for _, v := range words {
		if len(v) > 0 {
			wordMap[v]++
		}
	}
	for k, v := range wordMap {
		fmt.Println(k, ":", v)
	}
}

func helper(word string) ([]string, error) {
	re, err := regexp.Compile("[^a-zA-Z0-9']+")
	if err != nil {
		return nil, err
	}
	word = re.ReplaceAllString(word, " ")

	words := strings.Split(word, " ")

	return words, err
}
