package factory

import "fmt"

func GetDevice(deviceType string) (IDevice, error) {
	if deviceType == "smartphone" {
		return NewSmartphone(), nil
	}
	if deviceType == "laptop" {
		return NewLaptop(), nil
	}
	return nil, fmt.Errorf("wrong device type passed")
}

func PrintDetails(d IDevice) {
	fmt.Printf("Device: %s\n", d.GetName())
	fmt.Printf("Price: $%.2f\n", d.GetPrice())
	fmt.Println()
}
