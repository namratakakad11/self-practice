package main

import "fmt"

func main(){

	fmt.Println("executing")
	arr := [] int{1,2,3}
	ch := make(chan int, len(arr))
	go func(arr []int, ch chan int){
		for _,ele :=range arr{
			ch <- ele +9
		}
		close(ch)
		fmt.Println("after closing connection")
	}(arr,ch)

	for i:=0;i<len(arr);i++{
		fmt.Println(<-ch)
	}

}