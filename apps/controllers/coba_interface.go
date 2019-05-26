package controllers


type Human struct {
}

func (h Human) Eats() string {
	return "eats"
}

func (h Human) Drinsk() string {
	return "drinks"
}

func (h *Human) Death() string {
	return "death"
}
func NewHuman() *Human {
	human := &Human{}
	return human
}

