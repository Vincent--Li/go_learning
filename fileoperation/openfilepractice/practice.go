package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	fileName := "../test1.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Println("please enter the text to write")
	var line string
	for {
		fmt.Scanln(&line)
		fmt.Printf("your input is %s\n", line)
		if line == "n" {
			break
		} else {
			writer.WriteString(line)
			writer.Flush()
		}
	}
}
