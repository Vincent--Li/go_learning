package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("err", err)
	}
	defer file.Close()

	fmt.Printf("%v\n", file)

}
