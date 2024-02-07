package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("starting the program")
	arr :=[] int{2,3,4}

	ch :=make(chan int,len(arr))

	go demo(arr, ch)

	time.Sleep(time.Second) // we let go routine return all the values

	for i:=0;i<len(arr);i++{
		fmt.Println(<-ch)
		time.Sleep(time.Second)
	}

	// fmt.Println("updated result is ",result)



}

func demo(number []int, ch chan int){

	for _,ele := range number{
		ch <-ele +4
	}

}