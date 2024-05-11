package utils

import (
	"fmt"
	"strings"
)

func PrintAbCd(){
	for x := 'A'; x <= 'Z'; x++ {
		if x % 2 != 1 {
			fmt.Print(strings.ToLower(string(x)))
		} else {
			fmt.Print(string(x))
		}
	}
}
