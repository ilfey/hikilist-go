package options

import (
	"fmt"
	"reflect"
)

/*
InList checks if the value is in the list.

`list` can be a slice or a list of strings, int64 or float64.
If `list` is wrong type, it will return panic.

Works with string and numeric types.

Pointers are not supported.
*/
func InList(list any) Option {
	return func(v reflect.Value) (string, bool) {
		switch v.Kind() {
		case reflect.String:
			for _, item := range list.([]string) {
				if item == v.String() {
					return "", true
				}
			}

			return "Field \"%s\" must be in list " + fmt.Sprintf("%v", list), false
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			for _, item := range list.([]int64) {
				if item == v.Int() {
					return "", true
				}
			}

			return "Field \"%s\" must be in list " + fmt.Sprintf("%v", list), false

		case reflect.Float32, reflect.Float64:
			for _, item := range list.([]float64) {
				if item == v.Float() {
					return "", true
				}
			}

			return "Field \"%s\" must be in list " + fmt.Sprintf("%v", list), false
		}

		return "Field \"%s\" must be in list %v", false
	}
}
