package main

import (
	"fmt"
	"github.com/foldera/playground/interfaces/life"
)

func main() {
	model := new(life.Animal).With(65, 13, "male")
	fmt.Println(model)
	////////////////////////////
	cat := new(life.Cat).With("Gypsy", "Tamijan", model.Height, model.Weight, model.Gender())
	fmt.Println(cat)
	fmt.Println("cat says", life.Call(cat))
}
