package a

import "github.com/foldera/playground/import.cycle/b"

const PackageName = "A"

func A() string {
	return "Function a." + PackageName + "()"
}

type Object struct {
	Id uint
}

func (o *Object) From(b *b.Object) *Object {
	return &Object{b.Id}
}
