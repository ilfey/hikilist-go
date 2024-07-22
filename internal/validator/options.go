package validator

import (
	"fmt"
	"reflect"
)

// Option is a validator option. It returns true if the value is valid
type Option func(reflect.Value) (string, bool)

/*
NotNil checks if the value is not nil.

Works only with pointers.
*/
func NotNil() Option {
	return func(v reflect.Value) (string, bool) {
		if v.Kind() != reflect.Ptr {
			return "Field \"%s\" has invalid type", false
		}

		return "Field \"%s\" is required", !v.IsNil()
	}
}

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
		case reflect.Float32, reflect.Float64:
			return "Field \"%s\" is required", v.Float() != 0
		case reflect.Bool:
			return "Field \"%s\" is required", v.Bool()
		default:
			return "Field \"%s\" has invalid type", false
		}
	}
}

/*
LessThat checks if the value is less than `num`.

Works with all numeric types.
If the value is a pointer,
it checks if the pointer is not nil,
then checks if the value is less than `num`.
*/
func LessThat[T int64 | float64](num T) Option {
	return func(v reflect.Value) (string, bool) {
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return "Field \"%s\" must be less that " + fmt.Sprint(num), v.Int() < int64(num)
		case reflect.Float32, reflect.Float64:
			return "Field \"%s\" must be less that " + fmt.Sprint(num), v.Float() < float64(num)
		case reflect.Ptr:
			if !v.IsNil() {
				return LessThat(num)(v.Elem())
			}

			return "Field \"%s\" must be not null", false

		default:
			return "Field \"%s\" has invalid type", false
		}
	}
}

/*
GreaterThat checks if the value is greater than `num`.

Works with all numeric types.
If the value is a pointer,
it checks if the pointer is not nil,
then checks if the value is greater than `num`.
*/
func GreaterThat[T int64 | float64](num T) Option {
	return func(v reflect.Value) (string, bool) {
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return "Field \"%s\" must be greater that " + fmt.Sprint(num), v.Int() > int64(num)
		case reflect.Float32, reflect.Float64:
			return "Field \"%s\" must be greater that " + fmt.Sprint(num), v.Float() > float64(num)
		case reflect.Ptr:
			if !v.IsNil() {
				return GreaterThat(num)(v.Elem())
			}

			return "Field \"%s\" must be not null", false

		default:
			return "Field \"%s\" has invalid type", false
		}
	}
}

/*
LenLessThat checks if the length of the value is less than `num`.

Works with all collection types.
If the value is a pointer,
it checks if the pointer is not nil,
then checks if the length of the value is less than `num`.
*/
func LenLessThat(num int) Option {
	return func(v reflect.Value) (string, bool) {
		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
			return "Field \"%s\" must be less that " + fmt.Sprint(num), v.Len() < num
		case reflect.Ptr:
			if !v.IsNil() {
				return LenLessThat(num)(v.Elem())
			}

			return "Field \"%s\" must be not null", false

		default:
			return "Field \"%s\" has invalid type", false
		}
	}
}

/*
LenGreaterThat checks if the length of the value is greater than `num`.

Works with all collection types.
If the value is a pointer,
it checks if the pointer is not nil,
then checks if the length of the value is greater than `num`.
*/
func LenGreaterThat(num int) Option {
	return func(v reflect.Value) (string, bool) {
		switch v.Kind() {
		case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
			return "Field \"%s\" must be greater that " + fmt.Sprint(num), v.Len() > num
		case reflect.Ptr:
			if !v.IsNil() {
				return LenGreaterThat(num)(v.Elem())
			}

			return "Field \"%s\" must be not null", false

		default:
			return "Field \"%s\" has invalid type", false
		}
	}
}

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
		default:
			return "Field \"%s\" must be in list %v", false
		}
	}
}
