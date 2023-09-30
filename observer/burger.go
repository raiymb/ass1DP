package observer

type Burger struct {
	Name string
}

func NewBurger(name string) *Burger {
	return &Burger{Name: name}
}

func (b *Burger) String() string {
	return b.Name
}
