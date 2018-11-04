package main

import (
	"fmt"
	"github.com/newfolder31/yurko/daemons"
	"github.com/newfolder31/yurko/infrastructures"
)

func main() {
	fmt.Println("Init application")

	//initialize session
	infrastructures.InMemorySession = infrastructures.NewSession()

	daemons.Run()
}
