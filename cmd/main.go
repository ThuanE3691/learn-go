package main

import (
	"log"

	"github.com/ThuanE3691/learn-go/cmd/api"
)

func main(){
	port := ":8080"
	server := api.NewAPIServer(port, nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}