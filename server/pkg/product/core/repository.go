package core

type ProductRepository interface {
	Create(data *Product) (string, error)
}
