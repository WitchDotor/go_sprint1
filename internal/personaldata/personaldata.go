package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\t Вес: %.2f\t Рост: %.2f\n", p.Name, p.Weight, p.Height)
}
