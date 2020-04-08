package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/AXR8ELW/data-pump-file/accountservice/internal/app/dbclient"
	"github.com/AXR8ELW/data-pump-file/accountservice/internal/app/model"
)

var DBClient dbclient.IBoltClient

func GetAccount(s string) {
	fmt.Println("inside get account")
	// Read the account struct BoltDB
	accounts, err := DBClient.QueryAccountFilter()

	// If err, return a 404
	if err != nil {
		log.Printf("Account not found %s " + s)
		return
	}

	createFile(accounts)
	readFile()
}
func createFile(accounts model.Accounts) {
	fmt.Println("inside create file")

	file, _ := json.MarshalIndent(accounts, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}

func readFile() {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
	fmt.Printf("\nError: %v", err)
}
