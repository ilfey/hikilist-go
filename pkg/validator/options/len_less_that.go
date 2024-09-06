package options

import (
	"fmt"
	"reflect"
)

/*
LenLessThan checks if the length of the value is less than `num`.

Works with all collection types.
If the value is a pointer,
it checks if the pointer is not nil,
then checks if the length of the value is less than `num`.
*/
func LenLessThan(num int) Option {
	return func(v reflect.Value) (string, bool) {
		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
			return "Field \"%s\" must be less that " + fmt.Sprint(num), v.Len() < num
		case reflect.Ptr:
			if !v.IsNil() {
				return LenLessThan(num)(v.Elem())
			}

			return "Field \"%s\" must be not null", false
		}

		return "Field \"%s\" has invalid type", false
	}
}
