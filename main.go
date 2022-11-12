package main

import (
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

////////////////////
// Main function ///
////////////////////

func main() {
	http.HandleFunc("/", HomeHandlerFunc)
	http.HandleFunc("/echo", EchoHandlerFunc)
	http.HandleFunc("/invert", InvertMatrixHandlerFunc)
	http.HandleFunc("/flatten", FlattenMatrixHandlerFunc )
	http.HandleFunc("/sum", SumMatrixHandlerFunc)
	http.HandleFunc("/multiply", MultiplyMatrixHandlerFunc)
	http.ListenAndServe(":8080", nil)
}
