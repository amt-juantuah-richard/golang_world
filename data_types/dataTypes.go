package main

import (
	"fmt"
	"reflect"
)

var print = fmt.Println

func main() {
	print(reflect.TypeOf(34))
	print(reflect.TypeOf("s"))
	print(reflect.TypeOf('s'))
	print(reflect.TypeOf(0.44))
	print(reflect.TypeOf(true))
	print(reflect.TypeOf(0.00000000000001))
	print(reflect.TypeOf(100000000000000))
}
