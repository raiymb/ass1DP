package strategy

type Burger struct {
	Ingredients []string
}

func (b *Burger) AddIngredient(ingredient string) {
	b.Ingredients = append(b.Ingredients, ingredient)
}

func (b *Burger) GetIngredients() []string {
	return b.Ingredients
}
