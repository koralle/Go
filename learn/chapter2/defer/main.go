package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("./map/main.go")
	if err != nil {
		fmt.Println("Error!")
	}

	defer file.Close()
	fmt.Println("Closed.")
}