package entity

type ExchangeRepositoryInterface interface {
	Save(exchange *Exchange) error
}
