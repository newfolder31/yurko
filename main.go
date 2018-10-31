package main

import (
	"fmt"
	"daemons"
	"infrastructures"
)

func main() {
	fmt.Println("Init application")

	//initialize session
	infrastructures.InMemorySession = infrastructures.NewSession()

	daemons.Run()
}
