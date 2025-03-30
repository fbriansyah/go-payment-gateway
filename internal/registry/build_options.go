package registry

import (
	"fmt"
	"reflect"
)

type BuildOption func(v any) error

func ValidateImplements(checkV any) BuildOption {
	checkT := reflect.TypeOf(checkV)

	if checkT.Kind() == reflect.Ptr {
		checkT = checkT.Elem()
	}
	
	if checkT.Kind() != reflect.Interface {
		panic(fmt.Sprintf("%T is not an interface", checkV))
	}

	return func(v any) error {
		t := reflect.TypeOf(v)
		if !t.Implements(checkT) {
			return fmt.Errorf("%T does not implement %T", v, checkV)
		}

		return nil
	}
}