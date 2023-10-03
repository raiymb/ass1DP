package factory

type Laptop struct {
	Device
}

func NewLaptop() IDevice {
	return &Laptop{
		Device: Device{
			name:  "Laptop",
			price: 999.99,
		},
	}
}
