package b

import "github.com/foldera/playground/import.cycle/a"

const PackageName = "B"

func B() string {
	return "Function b." + PackageName + "()"
}

type Object struct {
	Id uint
}

func (o *Object) From(a *a.Object) *Object {
	return &Object{a.Id}
}
