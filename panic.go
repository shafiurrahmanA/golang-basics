package main

import (
	"errors"
	"fmt"
)

func A() {
	defer fmt.Println("its red signal stop the vehicle")
	B()
}
func B() {
	defer fmt.Println("And if it keeps getting hotter...")
	C()
}
func C() {
	defer fmt.Println("be quick only few seconds left..")
	Break()
}
func Break() {
	defer fmt.Println("Its green you can go...")
	panic(errors.New("or else you will be punished!!!"))

}
func main() {
	A()
}
