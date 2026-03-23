package models

type Storage interface{}

type Provider struct {
	storage Storage
}

func NewProvider(storage Storage) Provider {
	return Provider{
		storage: storage,
	}
}
