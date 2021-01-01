package main

import "fmt"

func main(){
	go func() {

		defer func(){
			if err := recover(); err != nil {
				fmt.Println("this is the error, ", err)
			}
		}()

		var hashMap map[int]string
		hashMap[0] = "2"
		fmt.Println("test")
	}()

	for i:=0 ; i< 100; i++ {
		fmt.Println(i)
	}

}
