package domain

type AccountRepository interface {
	Save(account *Account)
	FindByID(id string) (*Account, error)
	FindByApiKey(apiKey string) (*Account, error)
	Update(account *Account) error
}
