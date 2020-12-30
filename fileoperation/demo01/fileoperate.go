package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){

	file, err := os.Open("../test.txt")
	if err != nil {
		fmt.Println("err", err)
	}

	// 当函数退出时, 要及时关闭file
	defer file.Close()

	// 创建一个*Reader, 是带缓冲区的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')

		//输出内容
		fmt.Printf(str)

		if err == io.EOF {
			break
		}



	}

	fmt.Println("文件读取结束")

}
