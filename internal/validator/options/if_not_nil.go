package options

import "reflect"

/*
IfNotNil calls the next option if the value is not nil.

Works with pointers only.
*/
func IfNotNil(opts ...Option) Option {
	return func(v reflect.Value) (string, bool) {
		msg, ok := NotNil()(v)
		if !ok {
			return msg, false
		}

		for _, opt := range opts {
			msg, ok = opt(v)
			if !ok {
				return msg, false
			}
		}

		return "", true
	}
}
