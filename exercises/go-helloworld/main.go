package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sherlock Holmes sat silently...")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
