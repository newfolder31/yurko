package main

import (
	"fmt"
	"yurko/src/daemons"
)

func main() {
	fmt.Println("Init application")

	daemons.Run()
}
