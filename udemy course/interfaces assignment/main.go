package main

import (
	"fmt"
	"io"

	"os"
)



func main(){

	var argsWithoutProg = os.Args[1]
	fmt.Println(argsWithoutProg)

	fl, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, fl)

}