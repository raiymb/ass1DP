package abstractFactory

func UseChair(c IChair) {
	c.Sit()
}

func UseTable(t ITable) {
	t.Eat()
}
