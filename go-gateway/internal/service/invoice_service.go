package service

import (
	"github.com/igor-marchi/go-gateway/internal/domain"
	"github.com/igor-marchi/go-gateway/internal/dto"
)

type InvoiceService struct {
	repository     domain.InvoiceRepository
	accountService *AccountService
}

func NewInvoiceService(repository domain.InvoiceRepository, accountService *AccountService) *InvoiceService {
	return &InvoiceService{
		repository:     repository,
		accountService: accountService,
	}
}

func (s *InvoiceService) CreateInvoice(input dto.CreateInvoiceInput) (*dto.InvoiceOutput, error) {
	account, err := s.accountService.FindByApiKey(input.ApiKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, account.ID)
	if err != nil {
		return nil, err
	}

	err = invoice.Process()
	if err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(input.ApiKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	err = s.repository.Save(invoice)
	if err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil
}
