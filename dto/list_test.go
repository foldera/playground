package dto_test

import (
	"github.com/foldera/playground/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (Object) ToTitleProvider() dto.Mapper[Object, TitleProvider] {
	return func(from Object) (TitleProvider, error) {
		return TitleProvider{from.Title}, nil
	}
}

func TestNewList(t *testing.T) {
	models := []Object{{Title: "A1 title"}, {Title: "A2 title"}, {Title: "A3 title"}}
	list, err := dto.NewList[TitleProvider](models, Object{}.ToTitleProvider())
	if assert.Nil(t, err) {
		assert.Equal(t, 3, len(list))
	}
}
