package domain

import "bankappdemo/errs"

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
	// SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	// FindBy(accountId string) (*Account, *errs.AppError)
}
