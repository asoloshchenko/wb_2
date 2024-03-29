package pattern

import (
	"fmt"
	"math/rand"
)

/*
	Реализовать паттерн «фасад».

Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Facade_pattern
*/
type subsys1 struct {
	r rand.Rand
}

type subsys2 struct {
	delta int
}

func (s *subsys1) Method1() int {
	return s.r.Int()
}
func (s2 *subsys2) Method2(i int) int {
	fmt.Println("got:", i)
	return i + s2.delta
}

type facade struct {
	sys1 *subsys1
	sys2 *subsys2
}

func NewFacade(seed int, delta int) *facade {
	r := rand.New(rand.NewSource(int64(seed)))
	return &facade{
		sys1: &subsys1{r: *r},
		sys2: &subsys2{delta: delta},
	}
}

func (f *facade) Method() {
	i := f.sys1.Method1()
	j := f.sys2.Method2(i)
	println(i, j)
}
