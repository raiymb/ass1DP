package observer

type Observer interface {
	Update(*Burger)
}
