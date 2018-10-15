package main

import (
	"fmt"
	"yurko/daemons"
)

func main() {
	fmt.Println("Init application")

	daemons.Run()
}
