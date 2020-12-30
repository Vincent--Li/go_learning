package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){
	go func(){
		i := 0
		for {
			fmt.Println("hello,world from goroutine " + strconv.Itoa(i))
			i++
			time.Sleep(time.Second)
		}
	}()

	for i:=0; i< 10; i++ {

		fmt.Println("hello,world from main thread " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
