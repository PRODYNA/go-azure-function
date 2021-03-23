package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{ \"message\": \"Hello World!\" }"))
}

func main() {
	fmt.Println("START")

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		// function runs in azure
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/hello", helloHandler)
	fmt.Println("Listening on Port ", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
