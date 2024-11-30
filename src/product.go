package src

type ProductInterface interface {
	GetID() int
	GetName() string
	GetPrice() int
	GetWeight() int
}

type Product struct {
	Id     int
	Name   string
	Price  int
	Weight int
}

func (p *Product) GetID() int {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() int {
	return p.Price
}

func (p *Product) GetWeight() int {
	return p.Weight
}

func NewProduct(id int, name string, price int, weight int) ProductInterface {
	return &Product{
		Id:     id,
		Name:   name,
		Price:  price,
		Weight: weight,
	}
}
