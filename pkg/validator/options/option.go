package options

import "reflect"

// Option is a validator option. It returns true if the value is valid
type Option = func(reflect.Value) (string, bool)
