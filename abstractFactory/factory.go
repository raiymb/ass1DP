package abstractFactory

import "fmt"

func GetFurnitureFactory(brand string) (IFurnitureFactory, error) {
	if brand == "ikea" {
		return &IkeaFactory{}, nil
	}

	if brand == "ashley" {
		return &AshleyFactory{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
