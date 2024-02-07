package main

import(
	"fmt"
	"time"
)

func main(){

	n := 10000000

	testSlice  := [] int {} 

	ts2 := make([]int, 0, n)


	fmt.Println("without preallocation ", timeloop(testSlice,n))
	fmt.Println("with preallocation",timeloop(ts2,n))

}


func timeloop(slice []int, n int) time.Duration{
	var to = time.Now()

	for len(slice)<n{
		slice = append(slice, 1)
	}
	return time.Since(to)
}