package main

import (
	"log"
)

func main() {
	err := SendThing("localhost:10123")
	if err != nil {
		log.Fatal(err)
	}
}
