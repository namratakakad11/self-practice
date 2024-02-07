package main

import (
	"encoding/json"
	"fmt"
)

var jsonobj = `{
	"name" : "GOLinuxCloud",
"years_of_service" : "5",
"nature_of_company" : "Online Academy",
"no_of_staff" : "10"
}`

func Unmarshalling() {

	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonobj), &result)
	if err != nil {
		panic(err)
	}


	for k, v:=  range  result{
		fmt.Println(k,v)
	}

}
