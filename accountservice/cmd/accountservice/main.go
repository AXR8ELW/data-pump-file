package main

import (
	"fmt"

	"github.com/AXR8ELW/data-pump-file/accountservice/internal/app/dbclient"
	"github.com/AXR8ELW/data-pump-file/accountservice/internal/app/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient() // NEW
	fmt.Println("going to get account")
	service.GetAccount("10000")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.InitializeBucket()
	service.DBClient.SeedAccounts()
}
