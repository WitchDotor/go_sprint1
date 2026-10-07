package personaldata

type Personal struct {
	Name string,
	Weight float64б
	Height float64,
}

func (p Personal) Print() {
	Printf("Имя: %s\t Вес: %f\t Рост: %f")
}
