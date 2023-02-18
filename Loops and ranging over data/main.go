package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
	animals := []string{"dog", "fish", "cat", "cow"}
	for _, animal := range animals {
		log.Println(animal)
	}
	food := make(map[string]string)
	food["today"] = "fish"
	food["tomorrow"] = "hotdog"
	for x, i := range food {
		log.Println(x, i)
	}
}
