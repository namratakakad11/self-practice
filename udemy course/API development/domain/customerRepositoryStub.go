package domain

type CustomerRepositoryStub struct {
	Customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	Customer := []Customer{
		{"2312", "nsdjn", "dsfsd", "fdsrf", "fvds", "sdfwed"},
		{"2312", "nsdjn", "dsfsd", "fdsrf", "fvds", "sdfwed"},
		{"2312", "nsdjn", "dsfsd", "fdsrf", "fvds", "sdfwed"},
	}
	return CustomerRepositoryStub{Customer}
}
