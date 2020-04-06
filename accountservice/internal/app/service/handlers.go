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
	account, err := DBClient.QueryAccount(s)

	// If err, return a 404
	if err != nil {
		log.Printf("Account not found %s " + s)
		return
	}

	log.Printf("Account found %s", account.Name)
	createFile(account)
	readFile()
}
func createFile(account model.Account) {
	// file, err := os.Create("test.txt") // Truncates if file already exists, be careful!
	// if err != nil {
	// 	log.Fatalf("failed creating file: %s", err)
	// }
	// defer file.Close() // Make sure to close the file when you're done

	// len, err := file.WriteString("The Go Programming Language, also commonly referred to as Golang, is a general-purpose programming language, developed by a team at Google.")

	// if err != nil {
	// 	log.Fatalf("failed writing to file: %s", err)
	// }
	fmt.Println("inside create file")
	file, _ := json.MarshalIndent(account, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
	//fmt.Printf("\nLength: %d bytes", len)
	fmt.Println("\nFile Name: ")
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
