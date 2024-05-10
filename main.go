package main

import (
	"fmt"
	"war/utils"
)
func main(){
	var res []int
	for i := 1 ; i <= 50; i++ {
		if utils.IsPrime(i) {
			res = append(res, i)
			fmt.Println(res)
			
			
		}
	}
}