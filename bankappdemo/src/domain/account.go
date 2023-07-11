package domain

import (
	"bankappdemo/dtos"
	"bankappdemo/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

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

	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)

	FindBy(accountId string) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() *dtos.NewAccountResponse {
	return &dtos.NewAccountResponse{a.AccountId}
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: dbTSLayout,
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}
