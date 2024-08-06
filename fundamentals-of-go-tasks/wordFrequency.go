package main

import (
	"strings"
)

func WordFrequency(s string) map[string]int {
	words := strings.Fields(s)
	wordFrequency := make(map[string]int)
	for _, word := range words {
		wordFrequency[word]++
	}
	return wordFrequency
}