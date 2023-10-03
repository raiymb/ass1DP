package factory

type Device struct {
	name  string
	price float64
}

func (d *Device) SetName(name string) {
	d.name = name
}

func (d *Device) GetName() string {
	return d.name
}

func (d *Device) SetPrice(price float64) {
	d.price = price
}

func (d *Device) GetPrice() float64 {
	return d.price
}
