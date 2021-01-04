package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err= ", err)
			continue
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn Write err= ", err)
		}
		if line == "n\r\n" {
			return
		}

	}

}
