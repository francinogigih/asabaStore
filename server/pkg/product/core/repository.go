package core

type ProductRepository interface {
	Create(data *Product) (string, error)
	Update(data *Product, code string) error
	GetList(active bool) ([]Product, error)
}
