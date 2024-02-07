package main

import "fmt"

type ContInfo struct{
	email string
	pincode int
}

type Person struct{
	firstName string
	lastName string
	ContInfo
}

func main(){

	alex := Person{
		firstName:"Alex",
		lastName:"andrison",
		ContInfo: ContInfo{
			email: "alex@gmail.com",
			pincode: 123456,
		},
	}


	alex.print()



}

func (p Person) print(){
	fmt.Printf("%+v\n", p)
}