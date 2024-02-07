package main

import (
	"encoding/json"

	"os"
)

var saveJSON = "saveJSON.json"

type company struct {
	Name    string
	Address string
	Desc    string
}

func marhsallingindent() {

	companydata := []company{
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

	jsondata, err := json.MarshalIndent(companydata, "", "")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(saveJSON, jsondata, 0644)
	if err != nil {
		panic(err)
	}
}
