package observer

import "sync"

type BurgerOrder struct {
	mutex     sync.Mutex
	observers []Observer
	burgers   []*Burger
}

func NewBurgerOrder() *BurgerOrder {
	return &BurgerOrder{}
}

func (bo *BurgerOrder) Attach(observer Observer) {
	bo.mutex.Lock()
	defer bo.mutex.Unlock()
	bo.observers = append(bo.observers, observer)
}

func (bo *BurgerOrder) Detach(observer Observer) {
	bo.mutex.Lock()
	defer bo.mutex.Unlock()
	for i, obs := range bo.observers {
		if obs == observer {
			bo.observers = append(bo.observers[:i], bo.observers[i+1:]...)
			break
		}
	}
}

func (bo *BurgerOrder) Notify() {
	bo.mutex.Lock()
	defer bo.mutex.Unlock()
	for _, observer := range bo.observers {
		observer.Update(bo.burgers[len(bo.burgers)-1])
	}
}

func (bo *BurgerOrder) AddBurger(burger *Burger) {
	bo.mutex.Lock()
	bo.burgers = append(bo.burgers, burger)
	bo.mutex.Unlock()
	bo.Notify()
}
