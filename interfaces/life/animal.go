package life

import "fmt"

type Animal struct {
	Height float32 // this is a field
	Weight float32
	gender string
}

func (m *Animal) With(HeightCm, WeightKg float32, gender string) *Animal {
	m.Height = HeightCm
	m.Weight = WeightKg
	m.gender = gender
	return m
}

// String implements the fmt.Stringer interface for life.Animal
func (m *Animal) String() string {
	return fmt.Sprintf("Animal >>> Height:%0.2f, Weight: %0.2f, Gender: %s", m.Height, m.Weight, m.gender)
}

// Gender is a method/behaviour that returns private gender of Animal
func (m *Animal) Gender() string {
	return m.gender
}
