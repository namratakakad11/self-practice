package main

import (
	"fmt"
	"time"
)

func main() {

	// start :=time.Now()
	// defer func(){
	// 	fmt.Println(time.Since(start))
	// }()
	// evilNinjas := []string{"tommy","johnny","bobby","andy"}

	// for _, evilNinjas := range evilNinjas{
	// 	go attack(evilNinjas)
	// }

	// time.Sleep(time.Second *2)
	fmt.Println("executing..")
	go test(6)

	fmt.Println("done")
	time.Sleep(time.Second * 2)

}

// func attack(target string) {
// 	fmt.Println(target)
// 	time.Sleep(time.Second)
// }

func test(number int) {
	fmt.Println(number * 3)
}
