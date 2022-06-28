package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name = "Mateus"
	var version = 1.1
	var age = 28
	fmt.Println("Hello Mr.", name, "your age is", age)
	fmt.Println("This program is on version", version)

	fmt.Println("The type of the variable age is", reflect.TypeOf(age))
	fmt.Println("The type of the variable version is", reflect.TypeOf(version))
}
