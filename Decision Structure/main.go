package main

import "log"

func main() {
	var isTrue bool
	isTrue = true
	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
	// && || if...else if...else
	myVar := "cat"
	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	default:
		log.Println("cat is something else")
	}

}
