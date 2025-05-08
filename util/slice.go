package util

import "reflect"

func MakeAnySlice(input interface{}) []any {
	v := reflect.ValueOf(input)


	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}


	if v.Kind() != reflect.Slice {
		panic("MakeAnySlice: input must be slice or pointer to slice")
	}

	result := make([]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = v.Index(i).Interface()
	}

	return result
}


func MakeAssociationSlice[T any](associationName string, model []T) map[string][]any {
	return map[string][]any{
		associationName: MakeAnySlice(model),
	}
}

