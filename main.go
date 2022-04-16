package main

import "C"

import (
	"os"
)

//export helloWorld
func helloWorld() string {
	return "Hello World"
}

//export getUid
func getUid() int{
	return os.Getuid()
}


func main() {}
