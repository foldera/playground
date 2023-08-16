package composition

type ObjectA struct {
	ID    uint
	Title string
	Count uint32
}

type ObjectB struct {
	ID          uint
	Title       string
	Description string
}

type ObjectC struct {
	Value string
	Count uint64
}

type AmbiguousObject struct {
	ObjectA
	ObjectB
	*ObjectC
}

type ClearObject struct {
	ObjectA // Root
	B       ObjectB
	C       *ObjectC
}
