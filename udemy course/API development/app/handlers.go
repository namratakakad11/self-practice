package app

import (
	"APIDEVELOPMENT/service"
	"encoding/json"

	// "fmt"

	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomers()

	w.Header().Add("content-type", "application/json")

	json.NewEncoder(w).Encode(customers)

}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	cust, err := ch.service.GetCustomer(id)
	if err != nil {

		// to replace this create a new func write response
		// w.Header().Add("content-type", "application/json")
		// w.WriteHeader(err.Code)
		// // fmt.Fprintf(w, err.Message)
		// json.NewEncoder(w).Encode(err.Asmessage())

		WriteResponse(w, err.Code, err.Asmessage())
	} else {

		// to replace this create a new func write response
		// w.Header().Add("content-type", "application/json")
		// json.NewEncoder(w).Encode(cust)

		WriteResponse(w, http.StatusOK, cust)
	}

}

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	// fmt.Fprintf(w, err.Message)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)

	}
}

// func getCustomer(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)

// 	fmt.Fprint(w, vars["id"])

// }

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "post request recieved")
// }
