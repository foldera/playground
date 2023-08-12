package object

type Model interface {
	GetId() uint
	GetTitle() string
}

type Params struct {
	Id    *uint
	Title *string
}

type model struct {
	id    uint
	title string
}

func (m model) GetId() uint {
	return m.id
}

func (m model) GetTitle() string {
	return m.title
}

type Setter func(p Params) Model

func New() Setter {
	m := new(model)
	return func(p Params) Model {
		if p.Id != nil {
			m.id = *p.Id
		}
		if p.Title != nil {
			m.title = *p.Title
		}
		return m
	}
}
