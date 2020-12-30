package main

import (
	"fmt"
	"time"
)

func main(){
	intChan := make(chan int)
	exitChan := make(chan bool)
	loops := 10000

	go readData(intChan, exitChan, 1, 1)
	go readData(intChan, exitChan, 2,2)
	go readData(intChan, exitChan, 3,3)
	go writeData(intChan, exitChan, loops)

	<-exitChan
}

func writeData(intChan chan int, exitChan chan bool, loops int){
	for i:= 1; i <= loops ; i++ {
		intChan <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool, worker int, workload int){
	for item := range intChan {
		fmt.Printf("worker %d , handling %d\n", worker, item)
		time.Sleep(time.Second * time.Duration(workload))
	}

	exitChan<-true
	close(exitChan)
}

