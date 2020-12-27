package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {

	//var s, sep string
	//for i:=1; i< len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//fmt.Println(s)

	//for index, value := range os.Args[1:]{
	//	fmt.Println(index)
	//	fmt.Println(value)
	//}

	//fmt.Println(strings.Join(os.Args[1:], " "))

	//version 4


	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n{
		fmt.Println()
	}
}
