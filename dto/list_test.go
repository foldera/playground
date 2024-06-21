package dto_test

import (
	"fmt"
	"github.com/foldera/playground/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListOutput struct {
	list any
	err  error
}
type ListTest[I, O any] struct {
	name     string
	from     []I
	mapper   dto.Mapper[I, O]
	expected ListOutput
}

var testIntToStringList = ListTest[int, string]{
	"from []int to []string", []int{1, 2, 3},
	func(from int) (string, error) { return fmt.Sprintf("string.%d", from), nil },
	ListOutput{dto.List[string]{"string.1", "string.2", "string.3"}, nil},
}

func TestNewList(t *testing.T) {
	got, err := dto.NewList[string](testIntToStringList.from, testIntToStringList.mapper)
	assert.Equal(t, testIntToStringList.expected.err, err)
	assert.Equal(t, 3, len(got))
	assert.Equal(t, testIntToStringList.expected.list, got)
}

func TestList_Append(t *testing.T) {
	list, err := dto.NewList[string](testIntToStringList.from, testIntToStringList.mapper)
	assert.Equal(t, testIntToStringList.expected.err, err)
	assert.Equal(t, len(testIntToStringList.from), len(list))

	got := list.Append("string.4", "string.5")
	assert.Equal(t, len(list)+2, len(got))

	var nilList dto.List[string]
	gotFromNilList := nilList.Append("string.0")
	assert.Equal(t, 1, len(gotFromNilList))
}

//func TestNewPtrList(t *testing.T) {
//var obj *Object
//models := []TitleProvider{{"title.0"}, {"title.1"}, {"title.3"}}
//list, err := dto.NewList[Object](models, TitleProvider{}.ToObject(obj))
//if assert.Nil(t, err) {
//assert.Equal(t, 3, len(list))
//for i := range list {
//assert.Equal(t, models[i].title, list[i].Title)
//}
//}
//	models := []Object{{Title: "A1 title"}, {Title: "A2 title"}, {Title: "A3 title"}}
//	list, err := dto.NewPtrList[TitleProvider](models, Object{}.ToTitleProvider())
//	if assert.Nil(t, err) {
//		assert.Equal(t, 3, len(list))
//		for i := range list {
//			assert.Equal(t, models[i].Title, list[i].title)
//		}
//	}
//}
