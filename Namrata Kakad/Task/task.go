package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Name string
}

func main() {

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
		arr = append(arr, file.Name())
		//   arr =append(arr, file.Name())
		//   keyvalmap := sliceToMap(arr)

		// //   data := &Data{Name: file.Name()}
		//   b, err := json.Marshal(keyvalmap)
		//   if err != nil {
		//     fmt.Println(err)
		//     return
	}
	fmt.Println(arr)
	keyvalmap := sliceToMap(arr)

	//   data := &Data{Name: file.Name()}
	b, err := json.Marshal(keyvalmap)
	if err != nil {
		fmt.Println(err)
		return
		fmt.Println(string(b))
		// fmt.Println(b)
		//   jsonData, err :=json.Marshal(file.Name())
		//   if err != nil {
		// 	fmt.Printf("could not marshal json: %s\n", err)
		// 	return
		// }

		// fmt.Printf("json data: %s\n", jsonData)
	}
	fmt.Println(arr)
	//    data := &Data{Name: arr}
	// b, err := json.Marshal(arr)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(string(b))

}
func sliceToMap(slice []string) map[int]string {
	keyval := make(map[int]string)
	count := 0
	for i := 0; i < len(slice); i++ {

		key := count + 1
		value := slice[i]
		keyval[key] = value
		fmt.Println(keyval, count)
	}

	return keyval

}
