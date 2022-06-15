package collection

import "reflect"

type facade struct {
	Real reflect.Value
}
