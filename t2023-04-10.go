package main

import (
	"fmt"
	"log"
)

func t20230410() {
	fmt.Println("hello world")
	// check that files in package can be accessed
	log.Printf("a simple check that files in package can be accessed")
	fileOrgCheck()

	// returning a function type from a function
	log.Printf("a simple example of a function returning a function")
	areaF := getAreaFunc()
	res := areaF(2, 4)
	fmt.Println(res)

	// another example
	log.Printf("another example of a function returning a function")
	getName := getNameFunc()
	fmt.Println(getName())
}
func getAreaFunc() func(int, int) int {
	return func(x, y int) int {
		return x * y
	}
}
func getNameFunc() func() string {
	return func() string { return "Joe" }
}
