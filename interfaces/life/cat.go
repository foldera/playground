package life

import "fmt"

type Cat struct {
	Animal  // cat inherits life.Animal. life.Animal is the parent of life.Cat
	Name    string
	Address string
}

func (c *Cat) With(name, address string, heightCm, wightKg float32, gender string) *Cat {
	c.Name = name
	c.Address = address
	c.Animal = *new(Animal).With(heightCm, wightKg, gender)
	return c
}

func (c *Cat) makeSound() string {
	return "miaow!"

}

func (c *Cat) String() string {
	return fmt.Sprintf("%s, Cat >>> Name: %s, Address: %s", c.Animal.String(), c.Name, c.Address)
}
