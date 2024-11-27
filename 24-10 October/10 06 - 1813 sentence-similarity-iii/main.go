package main

import (
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	lStrWords := strings.Split(sentence1, " ")
	sStrWords := strings.Split(sentence2, " ")
	result := false
	for {
		if len(lStrWords) == 0 || len(sStrWords) == 0 {
			result = true
			break
		}
		if lStrWords[0] == sStrWords[0] {
			lStrWords = lStrWords[1:]
			sStrWords = sStrWords[1:]
			continue
		}
		if lStrWords[len(lStrWords)-1] == sStrWords[len(sStrWords)-1] {
			lStrWords = lStrWords[:len(lStrWords)-1]
			sStrWords = sStrWords[:len(sStrWords)-1]
			continue
		}
		break
	}
	return result
}
