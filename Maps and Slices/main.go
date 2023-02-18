package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	//map
	myMap := make(map[string]string)

	myMap["pet"] = "dog"
	log.Println(myMap["pet"])
	myMap1 := make(map[int]string)
	myMap1[1] = "a"
	log.Println(myMap1[1])

	me := User{
		FirstName: "Junyan",
		LastName:  "Weng",
	}
	myMap2 := make(map[string]User)
	myMap2["me"] = me
	log.Println(myMap2["me"].FirstName)

	//slice
	var mySlice []int
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 2)

	sort.Ints(mySlice)
	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	log.Println(numbers)

	log.Println(numbers[3:5])
}
