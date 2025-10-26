package main

import (
	"fmt"
	"net/http"
)
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, CryptoStats!")
}
func main() {
	fmt.Println("server up & running")
    http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":8080", nil)
}								