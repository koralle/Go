package main

import (
	"fmt"
)

func sum(nums ...int) (result int){
	for _, n := range nums {
		result += n
	}
	return result
}

func main(){
	fmt.Println(sum(1, 2, 3, 4))
}