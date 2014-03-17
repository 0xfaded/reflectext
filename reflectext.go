// +build reflectext

package reflectext

import (
	"reflect"
)

var Available = true

func FuncOf(in, out []reflect.Type, dotdotdot bool) reflect.Type {
	return reflect.FuncOf(in, out, dotdotdot)
}

func InterfaceOf(methods []reflect.Method) reflect.Type {
	return reflect.InterfaceOf(methods)
}

func StructOf(fields []reflect.StructField) reflect.Type {
	return reflect.StructOf(fields)
}

func ArrayOf(count int, elem reflect.Type) reflect.Type {
	return reflect.ArrayOf(count, elem)
}

func Name(t reflect.Type, pkgPath, name string) reflect.Type {
	return reflect.Name(t, pkgPath, name)
}
