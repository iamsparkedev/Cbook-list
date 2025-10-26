package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server up & running")
    
	http.ListenAndServe(":8080", nil)
}								