package options

import "reflect"

/*
Required checks if the value is not empty.

Works with all types.
If the value is a pointer,
it checks if the pointer is not nil,
then checks if the value is not empty.
*/
func Required() Option {
	return func(v reflect.Value) (string, bool) {
		switch v.Kind() {
		case reflect.Ptr:
			if !v.IsNil() {
				return Required()(v.Elem())
			}

			return "Field \"%s\" is required", false
		case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
			return "Field \"%s\" is required", v.Len() != 0
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return "Field \"%s\" is required", v.Int() != 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return "Field \"%s\" is required", v.Uint() != 0
		case reflect.Float32, reflect.Float64:
			return "Field \"%s\" is required", v.Float() != 0
		case reflect.Bool:
			return "Field \"%s\" is required", v.Bool()
		}

		return "Field \"%s\" has invalid type", false
	}
}
