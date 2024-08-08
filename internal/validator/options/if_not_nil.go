package options

import "reflect"

/*
IfNotNil calls the next option if the value is not nil.

Works with pointers only.
*/
func IfNotNil(opts ...Option) Option {
	return func(v reflect.Value) (string, bool) {
		if v.Kind() != reflect.Ptr {
			return "Field \"%s\" has invalid type", false
		}

		if !v.IsNil() {
			for _, opt := range opts {
				msg, ok := opt(v)
				if !ok {
					return msg, false
				}
			}
		}

		return "", true
	}
}
