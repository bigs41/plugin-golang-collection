package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sort(t *testing.T) {
	arr := []testModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	var res []int
	Chain(arr).Sort(func(n testModel, _ int) int {
		return n.ID
	}).Map(func(r testModel, _ int) int {
		return r.ID
	}).Value(&res)
	assert.Len(
		t,
		res,
		len(arr),
	)
	assert.EqualValues(
		t,
		res,
		[]int{1, 2, 3},
	)
}

func Test_SortBy(t *testing.T) {
	arr := []testModel{
		{ID: 2, Name: "two"},
		{ID: 1, Name: "one"},
		{ID: 3, Name: "three"},
	}
	var res []string
	Chain(arr).SortBy("id").Map(func(r testModel, _ int) string {
		return r.Name
	}).Value(&res)
	assert.Len(t, res, 3)
	assert.EqualValues(
		t,
		res,
		[]string{"one", "two", "three"},
	)
}
