package main

import "fmt"

type bot interface{
	getGreating() string
}

type engbot struct{}
type spanishbot struct{}

func (engbot) getGreating()string{
	return "Hello!"
}

func (spanishbot) getGreating() string{
	return "hola!"
}

func printGreating(b bot){
	fmt.Println(b.getGreating())
}
func main(){

	eb := engbot{}
	sb  := spanishbot{}

	printGreating(eb)
	printGreating(sb)

}