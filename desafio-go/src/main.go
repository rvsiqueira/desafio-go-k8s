package main

import (
	"fmt"
	"log"
	"net/http"
)

func grettings() string {
	return "Code.education Rocks!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><b>%s</b></html>", grettings())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
