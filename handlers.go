package main

import (
	"fmt"
	"net/http"
)



const Description = `
Welcome, this basic web server exposes the following endpoints

/echo
/invert
/flatten
/sum
/multiply

You must send a csv file containing a square matrix, For example
'''
1,2,3
4,5,6
7,8,9
'''

Here is an example using  Curl:
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
`

func HomeHandlerFunc(w http.ResponseWriter, r *http.Request) {			
	fmt.Fprint(w, Description)
}


func EchoHandlerFunc(w http.ResponseWriter, r *http.Request) {		
	records, err := extractMatrix(w, r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
		return	
	}
	
	// response := Stringify(records)
	response := StringifyMatrix(records)
	fmt.Fprint(w, response)
}


func InvertMatrixHandlerFunc(w http.ResponseWriter, r *http.Request) {		
	records, err := extractMatrix(w, r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
		return	
	}
	
	response := InvertMatrix(records)
	fmt.Fprint(w, response)
}


func FlattenMatrixHandlerFunc(w http.ResponseWriter, r *http.Request) {		
	records, err := extractMatrix(w, r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
		return	
	}
	
	response := FlatMatrix(records)
	fmt.Fprint(w, response, "\n")
}


func SumMatrixHandlerFunc(w http.ResponseWriter, r *http.Request) {		
	records, err := extractMatrix(w, r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
		return	
	}
	
	response, err := SumMatrix(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
		return
	}

	fmt.Fprint(w, response, "\n")
}


func MultiplyMatrixHandlerFunc(w http.ResponseWriter, r *http.Request) {		
	records, err := extractMatrix(w, r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
		return	
	}
	
	response, err := MultiplyMatrix(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
		return
	}

	fmt.Fprint(w, response, "\n")
}