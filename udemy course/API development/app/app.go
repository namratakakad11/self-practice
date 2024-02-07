package app

import (
	"APIDEVELOPMENT/domain"
	"APIDEVELOPMENT/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRespositoryDB())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id}", ch.getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customers/{id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
	
}
