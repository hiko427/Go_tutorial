package main

import "log"

func main() {
	var a string
	a = "1"

	log.Println("a is set to", a)
	pointFunc(&a)
	log.Println("a is set to", a)

}

func pointFunc(x *string) {
	log.Println(x)
	c := "2"
	*x = c
}
