package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/logicmason/cdig/handlers/rest"
)

func main() {

	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
