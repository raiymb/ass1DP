package factory

type Smartphone struct {
	Device
}

func NewSmartphone() IDevice {
	return &Smartphone{
		Device: Device{
			name:  "Smartphone",
			price: 499.99,
		},
	}
}
