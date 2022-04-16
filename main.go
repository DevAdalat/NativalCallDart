package main

import "C"

//export HelloWorld
func HelloWorld() string {
	return "Hello World"
}

//export Devide
func Devide(a, b int) int {
	return a / b
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {}
