//+build !reflectext

package reflectext

import (
	"reflect"
)

var Available = false

func FuncOf(in, out []reflect.Type, dotdotdot bool) reflect.Type {
	panic("reflect.FuncOf not avaliable, see https://code.google.com/r/carlchatfield-go-arrayof-structof")
}

func InterfaceOf(methods []reflect.Method) reflect.Type {
	panic("reflect.InterfaceOf not avaliable, see https://code.google.com/r/carlchatfield-go-arrayof-structof")
}

func StructOf(fields []reflect.StructField) reflect.Type {
	panic("reflect.StructOf not avaliable, see https://code.google.com/r/carlchatfield-go-arrayof-structof")
}

func ArrayOf(count int, elem reflect.Type) reflect.Type {
	panic("reflect.ArrayOf not avaliable, see https://code.google.com/r/carlchatfield-go-arrayof-structof")
}

func Name(t reflect.Type, pkgPath, name string) reflect.Type {
	panic("reflect.Name not avaliable, see https://code.google.com/r/carlchatfield-go-arrayof-structof")
}
