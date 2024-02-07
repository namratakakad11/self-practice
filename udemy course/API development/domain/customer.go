package domain

import "APIDEVELOPMENT/errs"

type Customer struct {
	Id      string 		
	Name    string
	City    string
	Zipcode string
	DOB     string
	Status  string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer,*errs.AppError)
}

