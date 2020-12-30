package main

import "fmt"

func main(){
	queue := make(chan int)

	go func(){
		for i:= 20; i>=1 ; i-- {

				result := 1
				for j := 1; j<= i; j++ {
					result *= j
				}
				fmt.Printf("writing data to %d\n", result)
				queue <- result
		}
	}()


	for item := range queue {
		fmt.Printf("about to handle item %d\n", item)

		if item == 1 {
			break
		}
	}

}
