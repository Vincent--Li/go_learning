/*
 * @Author: your name
 * @Date: 2021-01-03 09:28:48
 * @LastEditTime: 2021-01-03 22:46:03
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go_learning\socket01\server\server.go
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server start to listening ...")
	// net.Listen("tcp", "0.0.0.0:8888")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err= ", err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
			fmt.Printf("Accept() suc con%v\n", conn)
			fmt.Println("client ip is ", conn.RemoteAddr().String())

			for {
				var data []byte = make([]byte, 512)
				n, err := conn.Read(data)
				if err != nil {
					fmt.Println("conn.Read err = ", err)
					break
				}
				if string(data[:n]) == "n\r\n" {
					fmt.Println("byebye")
					return
				}

				fmt.Printf("receive: %s", string(data[:n]))
			}
		}(conn)

	}

}
