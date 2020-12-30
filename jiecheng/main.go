package main

import "fmt"

func main(){
	const UPPER int = 100
	queue := make(chan int, 10)
	defer close(queue)

	queue <- 1

	for item := range queue {
		fmt.Printf("about to handle item %d\n", item)

		if item == UPPER {
			break
		}

		go func(){
			facVal := 1
			for i := 1; i <= item ; i++  {
				facVal *= i
			}
			fmt.Printf("handle item %d, result is %v\n", item, facVal)
			queue <- item + 1
		}()
	}

}
