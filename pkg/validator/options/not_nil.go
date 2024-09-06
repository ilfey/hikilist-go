package options

import "reflect"

/*
NotNil checks if the value is not nil.

Works with pointers only.
*/
func NotNil() Option {
	return func(v reflect.Value) (string, bool) {
		if v.Kind() != reflect.Ptr {
			return "Field \"%s\" has invalid type", false
		}

		return "Field \"%s\" is required", !v.IsNil()
	}
}
