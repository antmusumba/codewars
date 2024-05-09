package utils

import (
	
	"strings"
	
)

func ToCamelCase(s string) string {
	var res strings.Builder
	words := strings.FieldsFunc(s, func(c rune) bool {
		return c == '-' || c == '_'
	})
	for i, word := range words {
		if i == 0 {
			res.WriteString(word)
		} else {
			res.WriteString(strings.Title(word))
		}
	}
	return res.String()
}