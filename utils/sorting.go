package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func Order(sentence string) string {
	if sentence == "" {
		return ""
	}
	words := strings.Split(sentence, " ")
	mySlice := make([]string, len(words))
	for _ , word := range words {
		for _, char := range word {
			if unicode.IsDigit(char) {
				index, _:= strconv.Atoi(string(char))
				mySlice[index-1] = word
				break

			}
		}
	}
	return strings.Join(mySlice, " ")
}