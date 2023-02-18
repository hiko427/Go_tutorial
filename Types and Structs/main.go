package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Junyan",
		LastName:  "Weng",
	}
	log.Println(user.FirstName)
	log.Println(user.LastName)
}
