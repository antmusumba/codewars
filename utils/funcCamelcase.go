package utils

import "strings"
func CamelCase(s string) string {
	var res strings.Builder
    mySlice := strings.Split(s ," ")
	for i, word := range mySlice {
		if i == 0 {
			res.WriteString(strings.Title(word))
		} else {
			res.WriteString(strings.Title(word))
		}
	}
	return res.String()
}