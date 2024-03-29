package domain

import (
	"bankappdemo/errs"
	"bankappdemo/logger"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	mySqlClient *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.mySqlClient.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.mySqlClient.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while Querying Customer Table. " + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	var customer Customer

	findByIdSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	err := d.mySqlClient.Get(&customer, findByIdSql, id)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while Scanning Customer Table. " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {

	// dbUser := os.Getenv("DB_USER")
	// dbPasswd := os.Getenv("DB_PASSWD")
	// dbAddr := os.Getenv("DB_ADDR")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	// dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	// client, err := sqlx.Open("mysql", dataSource)
	// if err != nil {
	// 	panic(err)
	// }

	// // See "Important settings" section.
	// client.SetConnMaxLifetime(time.Minute * 3)
	// client.SetMaxOpenConns(10)
	// client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{
		mySqlClient: dbClient,
	}
}
