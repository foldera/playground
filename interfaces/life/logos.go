package life

type Logos interface {
	makeSound() string
}

func Call(l Logos) string {
	if l == nil {
		return "nothing"
	}
	return l.makeSound()
}
