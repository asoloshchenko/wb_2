package pattern

/*
	Реализовать паттерн «строитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Builder_pattern
*/
type PizzaBuilder interface {
	setSize()
	setFillings()
	getPizza() *Pizaa
}

func NewPizzaBuilder(pizzaType string) PizzaBuilder {
	switch pizzaType {
	case "4cheese":
		return NewFourCheeseBuilder()
	case "pepperoni":
		return NewPepperoniBuilder()
	default:
		return nil
	}
}

type Pizaa struct {
	fillings []string
	size     string
}

type PepperoniBuilder struct {
	fillings []string
	size     string
}

func NewPepperoniBuilder() *PepperoniBuilder {
	return &PepperoniBuilder{}
}

func (p *PepperoniBuilder) setSize() {
	p.size = "small"
}

func (p *PepperoniBuilder) setFillings() {
	p.fillings = append(p.fillings, "pepperoni", "mozzarella")
}

func (p *PepperoniBuilder) getPizza() *Pizaa {
	return &Pizaa{fillings: p.fillings, size: p.size}
}

type FourCheeseBuilder struct {
	fillings []string
	size     string
}

func NewFourCheeseBuilder() *FourCheeseBuilder {
	return &FourCheeseBuilder{}
}

func (f *FourCheeseBuilder) setSize() {
	f.size = "large"
}

func (f *FourCheeseBuilder) setFillings() {
	f.fillings = append(f.fillings, "fontina", "mozzarella", "cheddar", "parmesan")
}

func (f *FourCheeseBuilder) getPizza() *Pizaa {
	return &Pizaa{fillings: f.fillings, size: f.size}
}
