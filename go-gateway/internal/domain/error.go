package domain

import "errors"

var (
	// ErrAccountNotFound é retornado quando uma conta não é encontrada
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplicateApiKey é retornado quando há tentativa de criar uma conta com uma chave de API duplicada
	ErrDuplicateApiKey = errors.New("api key already exists")
	// ErrInvoiceNotFound é retornado quando uma fatura não é encontrada
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorizedAccess é retornado quando há tentativa de acesso não autorizado a um recurso
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)
