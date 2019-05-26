package main

import (
	"demoGo/apps/controllers"
	"fmt"
)

type HumanInterface interface {
	Eats() string
	Drinsk() string
	Death() string
}



func main() {
	fmt.Printf(controllers.NewHuman().Death())
}

