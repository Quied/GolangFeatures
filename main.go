package main

import (
	"fmt"
	"reflect"
)

func what() func() int {
	value := 42

	return func() int {
		return value
	}
}

func MakeFunc(mssg string) {
	fmt.Println("Type: ", reflect.TypeOf(mssg))
}

func Temp(Amt ...int) {
	fmt.Println(len(Amt))
}

type Gola struct {
	One, Two string
}

func main() {

	temp := Gola{One: "First", Two: "Second"}

	fmt.Println(temp.One)
	fmt.Println(temp.Two)

	var masg string
	masg = "ODR"

	Temp(23, 32, 45, 23, 5, 23, 523)

	func() { fmt.Println("Lambda: ", masg) }()

	MakeFunc(masg)

	fmt.Println(reflect.TypeOf(masg))

}
