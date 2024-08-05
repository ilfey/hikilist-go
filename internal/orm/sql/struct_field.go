package sql

import (
	"fmt"
	"reflect"
)

type StructField struct {
	reflectValue reflect.Value
	name         string

	ignore bool
	idx    int
}

func (f *StructField) ReflectType() reflect.Type {
	return f.reflectValue.Type()
}

func (f *StructField) ReflectValue() reflect.Value {
	return f.reflectValue
}

func (f *StructField) Kind() reflect.Kind {
	return f.reflectValue.Kind()
}

func (f *StructField) Name() string {
	return f.name
}

func (f *StructField) Value() any {
	return f.reflectValue.Interface()
}

func (f *StructField) Snake() string {
	return toColumnName(f.Name())
}

func (f *StructField) GetColumn(table string) string {
	return fmt.Sprintf("%s.%s", table, f.Snake())
}

func (f *StructField) ValueString() string {
	value := f.Value()

	if f.reflectValue.Kind() == reflect.Pointer {
		if f.reflectValue.IsNil() {
			return "NULL"
		}

		return toString(f.reflectValue.Elem().Interface())
	}

	return toString(value)
}
