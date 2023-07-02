package services

import "bankappdemo/errs"

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	// MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}
