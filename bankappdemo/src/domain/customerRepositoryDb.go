package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	mySqlClient *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := d.mySqlClient.Query(findAllSql)

	if err != nil {
		log.Println("Error while Querying Customer Table. ", err.Error())

		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer

		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateofBirth, &customer.Status)

		if err != nil {
			log.Println("Error while Scanning Customer Table. ", err.Error())

			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:Sample@123$@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{
		mySqlClient: client,
	}
}
