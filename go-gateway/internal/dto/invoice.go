package dto

import (
	"time"

	"github.com/igor-marchi/go-gateway/internal/domain"
)

const (
	StatusPending           = string(domain.StatusPending)
	StatusApproved          = string(domain.StatusApproved)
	StatusRejected          = string(domain.StatusRejected)
	PaymentTypeCreditCard   = string(domain.PaymentTypeCreditCard)
	PaymentTypeDebitCard    = string(domain.PaymentTypeDebitCard)
	PaymentTypeBankTransfer = string(domain.PaymentTypeBankTransfer)
)

type CreateInvoiceInput struct {
	ApiKey         string
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentType    string  `json:"payment_type"`
	CardNumber     string  `json:"card_number"`
	CardHolderName string  `json:"card_holder_name"`
	CVV            string  `json:"cvv"`
	ExpiryMonth    int     `json:"expiry_month"`
	ExpiryYear     int     `json:"expiry_year"`
}

type InvoiceOutput struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoice(input CreateInvoiceInput, accountId string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		Number:         input.CardNumber,
		CardHolderName: input.CardHolderName,
		CVV:            input.CVV,
		ExpiryMonth:    input.ExpiryMonth,
		ExpiryYear:     input.ExpiryYear,
	}
	return domain.NewInvoice(
		accountId,
		input.Description,
		input.Amount,
		domain.PaymentType(input.PaymentType),
		card,
	)
}

func FromInvoice(invoice *domain.Invoice) *InvoiceOutput {
	return &InvoiceOutput{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Amount:         invoice.Amount,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentType:    string(invoice.PaymentType),
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
