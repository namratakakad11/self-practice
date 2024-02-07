package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi, from slices")

	myslice := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("type of %T\n", myslice)

	myslice[0] = 3

	myslice[5] = 23

	myslice = append(myslice, 22, 33, 44, 55, 55)

	fmt.Print(myslice)

	newslices := append(myslice[5:5])
	fmt.Print(newslices)

	num := make([]int, 4)

	fmt.Printf("type of %T\n", num)

	fmt.Println("lenght ", len(num))

	num[0] = 1
	num[1] = 2
	num[2] = 1
	num[3] = 3
	// num[4] = 4

	num = append(num, 5, 6, 7, 8)
	fmt.Println(num)

	//  remove  values from slices using index

	index := 4

	myslice = append(myslice[:index], myslice[index+1:]...)

	fmt.Println("my sclice", myslice)

	mystring  := "abcdefgh %"

	fmt.Println(len(mystring))

	fmt.Println(len("hello world"))


	// datatypes

	var flo float32 = 102.3

	var num1 int32 = 36

	var c  = flo + float32(num1)
	fmt.Println("",c)





	

}
