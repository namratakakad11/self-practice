package main

import "fmt"

func main() {
	fmt.Println("starting program")
	ch := make(chan int) // make channel to transfer data between go routine

	go test1(5, ch)

	result := <-ch
	fmt.Println(" result is ", result)

}

func test1(number int, ch chan int) {
	result := number * 3
	fmt.Println("   ", result)
	ch <- result

}
