package singletone

import (
	"sync"
)

type CoffeeMachine struct {
	IsBusy bool
}

var once sync.Once
var instance *CoffeeMachine

func GetCoffeeMachine() *CoffeeMachine {
	once.Do(func() {
		instance = &CoffeeMachine{IsBusy: false}
	})
	return instance
}
