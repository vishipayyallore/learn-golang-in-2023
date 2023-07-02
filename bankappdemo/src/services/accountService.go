package services

import (
	"bankappdemo/domain"
	"bankappdemo/dtos"
	"bankappdemo/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(request dtos.NewAccountRequest) (*dtos.NewAccountResponse, *errs.AppError)
	// MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dtos.NewAccountRequest) (*dtos.NewAccountResponse, *errs.AppError) {
	// if err := req.Validate(); err != nil {
	// 	return nil, err
	// }
	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	if newAccount, err := s.repo.Save(account); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDto(), nil
	}
}
