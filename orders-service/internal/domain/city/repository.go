package city

type CityRepository interface {
	Find(id CityID) (City, error)
	FindByName(name string) (City, error)
	Create(name string) (CityID, error)
	Delete(id CityID) error
}
