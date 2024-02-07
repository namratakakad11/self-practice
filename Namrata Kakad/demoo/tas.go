package main

import (
	"fmt"
	"os"
	"encoding/json"
)
type Data struct {
    Name string
	num interface{}
}

func main(){

	arr := []string{}
	dir, err := os.Open(".")
   if err != nil {
      fmt.Println("Error:", err)
      return
   }
   defer dir.Close()
   files, err := dir.Readdir(-1)
   if err != nil {
      fmt.Println("Error:", err)
      return
   }
   for _, file := range files {
      fmt.Println(file.Name())
	  arr =append(arr, file.Name())
	  
	  
   }
   b, err := json.Marshal(arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}