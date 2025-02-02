package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from Service 1")
}

func main() {
	http.HandleFunc("/api/", handler)
	fmt.Println("Service 1 running on port 8081")
	http.ListenAndServe(":8081", nil)
}
