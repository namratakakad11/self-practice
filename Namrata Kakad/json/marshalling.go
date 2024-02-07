package main

import (
	"encoding/json"
	"fmt"
	"os"
)
var saveJSON1 = "saveJSON.json"
type company1 struct {
	Name    string
	Address string
	Desc    string
}

func main(){

	companydata := []company1{
		{
			Name:    "abc",
			Address: "adsjfhg",
			Desc:    "csd",
		},
		{
			Name:    "saf",
			Address: "adsjfhg",
			Desc:    "csd",
		},
		{
			Name:    "deqw",
			Address: "adsjfhg",
			Desc:    "csd",
		},
	}


	jsondata, err := json.Marshal(companydata)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(jsondata)

	jsonoutput := string(jsondata)

	fmt.Println(jsonoutput)

	err = os.WriteFile(saveJSON1,jsondata,0644)
	if err != nil{
		fmt.Println(err)
	}

	Unmarshalling()

	marhsallingindent()
}