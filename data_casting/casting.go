package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var print = fmt.Println

func main() {
	numb := 56
	str := rune(numb)
	str2 := strconv.Itoa(numb)
	digit, err := strconv.Atoi(string(str))
	print(str, reflect.TypeOf(str))
	print(str2, reflect.TypeOf(str2))
	print(digit, err, reflect.TypeOf(digit))
}
