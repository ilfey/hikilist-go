package options

import (
	"fmt"
	"reflect"
)

/*
GreaterThan checks if the value is greater than `num`.

Works with all numeric types.
If the value is a pointer,
it checks if the pointer is not nil,
then checks if the value is greater than `num`.
*/
func GreaterThan[T int64 | float64 | uint64](num T) Option {
	return func(v reflect.Value) (string, bool) {
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return "Field \"%s\" must be greater that " + fmt.Sprint(num), v.Int() > int64(num)
		case reflect.Float32, reflect.Float64:
			return "Field \"%s\" must be greater that " + fmt.Sprint(num), v.Float() > float64(num)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return "Field \"%s\" must be greater that " + fmt.Sprint(num), v.Uint() > uint64(num)
		case reflect.Ptr:
			if !v.IsNil() {
				return GreaterThan(num)(v.Elem())
			}

			return "Field \"%s\" must be not null", false
		}

		return "Field \"%s\" has invalid type", false
	}
}
