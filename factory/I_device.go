package factory

type IDevice interface {
	SetName(name string)
	SetPrice(price float64)
	GetName() string
	GetPrice() float64
}
