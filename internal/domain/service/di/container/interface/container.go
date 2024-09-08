package containerInterface

import "reflect"

type Container interface {
	Set(service any, alias reflect.Type) Container
	Has(key reflect.Type) bool
	Get(key reflect.Type) (reflect.Value, error)
}
