package main

import (
	"datapipeline/data-pump-file/accountservice/internal/app/service"
	"fmt"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient() // NEW
	service.StartWebServer("6767")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
