package main

import (
	"fmt"

	"github.com/thanhpk/randstr"
)

func Randomstr() {

	// does generate random same string
	/*stringtoencode := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	encoded := base64.StdEncoding.EncodeToString([]byte(stringtoencode))
	fmt.Println(encoded)**/
	// Random()

	encoded := randstr.String(80)
	fmt.Println(encoded)
}
