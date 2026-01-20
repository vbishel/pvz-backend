package product

type ProductRepository interface {
	Find(id ProductID) (Product, error)
	FindByName(name string) (Product, error)
	Create(name string) (ProductID, error)
	Delete(id ProductID) error
}
