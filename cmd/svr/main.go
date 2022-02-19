package main

import (
	"fmt"
	"log"
)

func main() {
	defer deathScream()

	fmt.Println(" === hello ===")
}

func deathScream() {
	if r := recover(); r != nil {
		log.Println(fmt.Errorf("I panicked and am quitting: %v", r))
		log.Println("I should be alerting someone...")
	}
}
