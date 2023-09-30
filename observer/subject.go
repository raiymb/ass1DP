package observer

type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}
