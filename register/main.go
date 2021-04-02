package main

import (
	"fmt"
	"reflect"
)

func hello() {
	fmt.Println("Hello world!")
}

func main() {
	/*h1 := hello
	h1()*/

	h1 := hello
	fv := reflect.ValueOf(h1)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)

}
