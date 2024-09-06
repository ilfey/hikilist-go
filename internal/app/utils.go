package app

import "reflect"

func reflectTypeOfNil[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil))
}
