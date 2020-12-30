package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	const UPPER int = 100
	queue := make(chan int, 10)
	//defer close(queue)
	r := rand.New(rand.NewSource(99))

	go func(){
		for {
			time.Sleep(time.Millisecond)
			tmpInt := r.Int() + 1;
			fmt.Printf("sending message %d\n", tmpInt)
			queue <- tmpInt
			if tmpInt == UPPER {
				close(queue)
			}
		}
	}()

	go func(){
		for item := range queue {
			fmt.Printf("goroutine 1 to handle item %d\n", item)
		}
	}()

	go func(){
		for item := range queue {
			fmt.Printf("goroutine 2 to handle item %d\n", item)
		}
	}()

	go func(){
		for item := range queue {
			fmt.Printf("goroutine 3 to handle item %d\n", item)
		}
	}()

	for item := range queue {
		fmt.Printf("main thread to handle item %d\n", item)
	}
}
