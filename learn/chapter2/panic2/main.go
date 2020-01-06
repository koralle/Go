package main


import (
	"fmt"
	"log"
	"errors"
)

func main(){
	defer func(){
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()

	a := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		panic(errors.New("index out of range."))
		fmt.Println(a[i])
	}
}