package main

import "fmt"

func main(){
	a :=20
	b :=30
// fmt.Println(swap(&a,&b))
c, d := swap(&a,&b)
fmt.Println(*c,*d)
}

func swap(a,b *int)(*int,*int){

	*a= *a+*b

	*b=*a-*b

	*a=*a-*b
	return a,b

}