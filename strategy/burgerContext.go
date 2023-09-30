package strategy

type Context struct {
	builder BurgerBuilder
}

func NewBurgerContext(builder BurgerBuilder) *Context {
	return &Context{
		builder: builder,
	}
}

func (bc *Context) CreateBurger() *Burger {
	return bc.builder.BuildBurger()
}
