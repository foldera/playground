package main

import (
	"github.com/foldera/playground/import.cycle/a"
	"github.com/foldera/playground/import.cycle/b"
	"log"
)

func main() {
	x := a.A()
	y := b.B()
	log.Println(x, y)
}
