package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {

	//variable
	var app application
	app.Domain = "mydomain.com"

	//db connection

	//Routes


	//server

	log.Println("Starting Server on Port : ", port);
	
	// http.HandleFunc("/", Hello);
	
	err:= http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err!= nil {
        log.Fatal(err);
    }

}