package caller

import (
	"reflect"
)

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
