package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

//Port
var Port = ":5555"

func main() {

	ConnectDB()

	http.HandleFunc("/", ThatsAFunk)
	fmt.Println("Serving @ : ", "http://127.0.0.1"+Port)
	log.Fatal(http.ListenAndServe(Port, nil))
}
