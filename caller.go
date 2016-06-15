package caller

import (
	"reflect"
)

// CallWith calls the function fn with the given parameters, regardless of the function signature.
func CallWith(fn interface{}, parameters ...interface{}) {
	va := reflect.ValueOf(fn)
	if va.Kind() != reflect.Func {
		panic("Not a function")
	}

	s := make([]reflect.Value, 0)
	for _, p := range parameters {
		v := reflect.ValueOf(p)
		s = append(s, v)
	}

	va.Call(s)
}
