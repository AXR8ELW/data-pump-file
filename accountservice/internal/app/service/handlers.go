package service

import (
	"log"
	"net/http"

	"github.com/AXR8ELW/data-pump-file/accountservice/internal/app/dbclient"
)

var DBClient dbclient.IBoltClient

func StartWebServer(port string) {

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
