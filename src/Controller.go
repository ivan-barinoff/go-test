package main

import (
	"fmt"
	"net/http"
)

func printHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>%s</h1><div>%s</div>", "Header", "Printed")
}

func main() {
	http.HandleFunc("/print/", printHandler)
	http.ListenAndServe(":8080", nil)
}