package main

import (
	"fmt"
)

func main(){
	var month map[int]string = map[int]string{}
	month[1] = "January"
	month[2] = "February"
	fmt.Println(month)

	jan := month[1]
	fmt.Println(jan)

	//delete(month, 1)
	//fmt.Println(month)

	for key, value := range month {
		fmt.Printf("%d %s\n", key, value)
	}

}