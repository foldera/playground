package dto_test

import (
	"fmt"
	"github.com/foldera/playground/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	nilTitleProvider   *TitleProvider
	happyTitleProvider = TitleProvider{"this is a title from TitleProvider"}
)

type output struct {
	object *Object
	err    error
}
type Test[I, O any] struct {
	name      string
	from      I
	converter dto.Mapper[I, O]
	expected  output
}

func TestNew_FromInputToOutput(t *testing.T) {
	var obj *Object
	converter := TitleProvider{}.ToObject(obj)
	var tests = []Test[TitleProvider, Object]{
		{"nil converter", happyTitleProvider, nil, output{nil, fmt.Errorf("trying to create new dto_test.Object without any mapper")}},
		{"from empty", TitleProvider{}, converter, output{nil, fmt.Errorf("trying to create new dto_test.Object using zero value of dto_test.TitleProvider")}},
		{"happy test", happyTitleProvider, converter, output{&Object{Title: happyTitleProvider.title}, nil}},
	}
	for _, test := range tests {
		var got *Object
		var err error
		t.Run(test.name, func(t *testing.T) {
			got, err = dto.New[Object](test.from, test.converter)
			assert.Equal(t, test.expected.err, err)
			assert.Equal(t, test.expected.object, got)
		})
	}
}

func TestNew_FromPtrInputToOutput(t *testing.T) {
	var obj *Object
	converter := TitleProvider{}.PtrToObject(obj)
	var tests = []Test[*TitleProvider, Object]{
		{"nil converter", &happyTitleProvider, nil, output{nil, fmt.Errorf("trying to create new dto_test.Object without any mapper")}},
		{"from nil", nilTitleProvider, converter, output{nil, fmt.Errorf("trying to create new dto_test.Object using zero value of *dto_test.TitleProvider")}},
		{"happy test", &happyTitleProvider, converter, output{&Object{Title: happyTitleProvider.title}, nil}},
	}
	for _, test := range tests {
		var got *Object
		var err error
		t.Run(test.name, func(t *testing.T) {
			got, err = dto.New[Object](test.from, test.converter)
			assert.Equal(t, test.expected.err, err)
			assert.Equal(t, test.expected.object, got)
		})
	}
}

type (
	Object        struct{ Id, Title string }
	TitleProvider struct{ title string }
	IdProvider    struct{ id string }
)

func (TitleProvider) ToObject(o *Object) dto.Mapper[TitleProvider, Object] {
	return func(from TitleProvider) (Object, error) {
		if o == nil {
			return Object{Title: from.title}, nil
		}
		o.Title = from.title
		return *o, nil
	}
}

func (TitleProvider) PtrToObject(o *Object) dto.Mapper[*TitleProvider, Object] {
	return func(from *TitleProvider) (Object, error) {
		if o == nil {
			return Object{Title: from.title}, nil
		}
		o.Title = from.title
		return *o, nil
	}
}

func (IdProvider) ToObject() dto.Mapper[IdProvider, Object] {
	return func(from IdProvider) (Object, error) {
		return Object{Title: from.id}, nil
	}
}

func (IdProvider) PtrToObject() dto.Mapper[*IdProvider, Object] {
	return func(from *IdProvider) (Object, error) {
		return Object{Id: from.id}, nil
	}
}
