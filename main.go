package main

import (
	"fmt"
	"github.com/newfolder31/yurko/daemons"
	"github.com/newfolder31/yurko/infrastructures"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "folder"
	password = "folder"
	dbname   = "folder"
)

func main() {
	fmt.Println("Init application")

	//initialize session
	infrastructures.InMemorySession = infrastructures.NewSession()
	db := infrastructures.ConnectToDB(host, user, password, dbname, port)
	daemons.Run()
}
