package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", postHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
