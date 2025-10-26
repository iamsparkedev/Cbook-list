package main

import (
	"cbook-list/api"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server up & running")
    
	http.HandleFunc("/", api.HandleRoot)
	http.ListenAndServe(":8080", nil)
}								