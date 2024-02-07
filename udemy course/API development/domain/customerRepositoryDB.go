package domain

import (
	"APIDEVELOPMENT/errs"
	"APIDEVELOPMENT/logger"
	// "database/sql"

	// "log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (db CustomerRepositoryDB) FindAll() ([]Customer, error) {
	// var row *sql.DB
	// connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	// client, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	customers := make([]Customer, 0)

	findallsql := "select id, name, city, zipcode, dob, status from customers"
	err := db.client.Select(&customers, findallsql)

	// create table customers(id varchar(20),name varchar(20),city varchar(20),zipcode varchar(20),dob varchar(20), status varchar(20) );

	// insert into customers(id,name,city,zipcode,dob,status) values('sdffsd','fsdfsdf','sdfgsd','sfdssdg','dsgdsg','dfgerfr');

	// replaced by sqlxselect command
	// row, err = db.client.Query(findallsql)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// err = sqlx.StructScan(row, &customers)
	// if err != nil {
	// 	logger.Error(err.Error())
	// }
	// replaced by sqlx
	// for row.Next() {

	// 	var c Customer
	// 	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DOB, &c.Status)
	// 	if err != nil {
	// 		logger.Error(err.Error())
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil

}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customeridsql := "select id, name, city, zipcode, dob, status from customers where id = $1"
	row := d.client.QueryRow(customeridsql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DOB, &c.Status)
	if err != nil {
		// log.Fatal("errror : ", err)
		return nil, errs.NewNotFoundError("customer not found")
	}
	return &c, errs.Newunexpectederror("unexpected database error")
}

func NewCustomerRespositoryDB() CustomerRepositoryDB {

	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	client, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.Error(err.Error())
	}
	return CustomerRepositoryDB{client}
}
