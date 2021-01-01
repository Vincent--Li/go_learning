package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	NumberOfConcurrency int = 10
	UpperNumber int = 10000
)

func main(){
	inputChan := make(chan int, 2000)
	resultChan := make(chan string, 2000)
	exitChan := make(chan bool, NumberOfConcurrency)

	go putNumber(inputChan)

	for i:=0; i < NumberOfConcurrency; i++ {
		go judgePrime(inputChan, resultChan, exitChan, i)
	}

	go processExit(resultChan, exitChan)

	for item := range resultChan {
		fmt.Printf("%v\n",item)
	}

}

func processExit(resultChan chan string, exitChan chan bool) {
	for i:= 0; i < NumberOfConcurrency; i++ {
		<- exitChan
	}
	close(exitChan)
	close(resultChan)
}

func judgePrime(inputChan chan int, resultChan chan string, exitChan chan bool, number int) {
	for  {
		item, ok := <- inputChan
		if !ok {
			break
		}

		time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))

		flag := true
		for i:= 2; i<item ; i ++ {
			if item % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resultChan <- strconv.Itoa(number) + " " + strconv.Itoa(item)
		}
	}

	exitChan <- true
}

func putNumber(inputChan chan int) {
	defer close(inputChan)

	for i:=2; i< UpperNumber; i++ {
		inputChan <- i
	}
}
