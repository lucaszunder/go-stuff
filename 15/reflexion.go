package main

import "reflect"

func Iterate(x interface{}, fn func(incoming string)) {
	value := reflect.ValueOf(x)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
