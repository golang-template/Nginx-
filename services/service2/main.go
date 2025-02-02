package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from Service 2")
}

func main() {
	http.HandleFunc("/api/", handler)
	fmt.Println("Service 2 running on port 8082")
	http.ListenAndServe(":8082", nil)
}
