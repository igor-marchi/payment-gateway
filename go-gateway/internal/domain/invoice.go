package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string
type PaymentType string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

const (
	PaymentTypeCreditCard   PaymentType = "credit_card"
	PaymentTypeDebitCard    PaymentType = "debit_card"
	PaymentTypeBankTransfer PaymentType = "bank_transfer"
)

type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    PaymentType
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CardHolderName string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
}

func NewInvoice(accountID, description string, amount float64, paymentType PaymentType, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.Number[len(card.Number)-4:]

	invoice := &Invoice{
		ID:             uuid.New().String(),
		Amount:         amount,
		AccountID:      accountID,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return invoice, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 1000 {
		return nil
	}

	randomValue := rand.New(rand.NewSource(time.Now().Unix())).Float64()
	if randomValue < 0.7 {
		i.UpdateStatus(StatusApproved) // 70% de chance de aprovação
	} else {
		i.UpdateStatus(StatusRejected) // 30% de chance de rejeição
	}

	i.UpdatedAt = time.Now()
	return nil
}

func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != StatusPending {
		return ErrInvalidStatus
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
