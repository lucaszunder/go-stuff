package main

import "reflect"

func Iterate(x interface{}, fn func(incoming string)) {
	value := reflect.ValueOf(x)
	field := value.Field(0)
	fn(field.String())
}
