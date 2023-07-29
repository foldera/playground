package life

type Logos interface {
	makeSound() string
}

func Call(l Logos) string {
	return l.makeSound()
}
