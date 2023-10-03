package main

import (
	"ass1/decorator"
	"ass1/observer"
	"ass1/strategy"
	"fmt"
)

func main() {
	//Strategy pattern
	fmt.Println("Strategy pattern:")
	classicBuilder := &strategy.ClassicBurgerBuilder{}
	classicContext := strategy.NewBurgerContext(classicBuilder)
	classicBurger := classicContext.CreateBurger()

	fmt.Println("Classic Burger Ingredients:")
	fmt.Println(classicBurger.GetIngredients())

	veggieBuilder := &strategy.VeggieBurgerBuilder{}
	veggieContext := strategy.NewBurgerContext(veggieBuilder)
	veggieBurger := veggieContext.CreateBurger()

	fmt.Println("\nVeggie Burger Ingredients:")
	fmt.Println(veggieBurger.GetIngredients())
	fmt.Println()
	//Observer pattern
	fmt.Println("Observer pattern:")
	burgerOrder := observer.NewBurgerOrder()

	customer1 := observer.NewCustomer("Big Smoke")
	customer2 := observer.NewCustomer("Raiymbek")

	burgerOrder.Attach(customer1)
	burgerOrder.Attach(customer2)

	burger1 := observer.NewBurger("Cheeseburger")
	burger2 := observer.NewBurger("Chicken Burger")

	burgerOrder.AddBurger(burger1)
	burgerOrder.AddBurger(burger2)

	burgerOrder.Detach(customer2)

	burger3 := observer.NewBurger("Number 9 Large")
	burgerOrder.AddBurger(burger3)

	fmt.Println("Decorator:")
	coffee := &decorator.Espresso{}
	coffeeWithCappuccino := &decorator.Cappuccino{
		CoffeeDecorator: decorator.CoffeeDecorator{Coffee: coffee},
	}
	coffeeWithCappuccinoAndVanillaLatte := &decorator.VanillaLatte{
		CoffeeDecorator: decorator.CoffeeDecorator{Coffee: coffeeWithCappuccino},
	}
	fmt.Printf("Price of Espresso with Cappuccino and Vanilla Latte topping is %d\n", coffeeWithCappuccinoAndVanillaLatte.GetPrice())
}
