package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(rw http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/home" {

		http.Error(rw, "not found ", http.StatusNotFound)

	}

	if r.Method != "GET" {
		http.Error(rw, "method type not supported", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Hello!")

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse form : %v", err)
		return
	}
	fmt.Fprintf(w, "post request successfull")
	name := r.FormValue("name")
	add := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", add)

}

func main() {

	fmt.Println("startiing")

	filehandle := http.FileServer(http.Dir("./static"))
	http.Handle("/", filehandle)

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/form", formHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("byeee..")
	}

}
