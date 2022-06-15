package collectionn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reject(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	var res []int
	Chain(arr).Reject(func(n, i int) bool {
		return n%2 == 0
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]int{1, 3},
	)
}

func Test_RejectBy(t *testing.T) {
	arr := []testModel{
		{ID: 1, Name: "one"},
		{ID: 2, Name: "two"},
		{ID: 3, Name: "three"},
	}
	var res []testModel
	Chain(arr).RejectBy(map[string]interface{}{
		"Id": 1,
	}).Value(&res)
	assert.EqualValues(
		t,
		res,
		[]testModel{
			{ID: 2, Name: "two"},
			{ID: 3, Name: "three"},
		},
	)
}
