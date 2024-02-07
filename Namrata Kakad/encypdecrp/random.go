package main

import (
	"fmt"
	"math/rand"
	"time"
)


func  Random(){
	rand.NewSource(time.Now().UnixNano())

	fmt.Println(rand.Intn(100))
}
