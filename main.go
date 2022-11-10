package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

// Shared Utils //
func validatePayload(paylaod [][]string) error {
	if len(paylaod) == 0 {
		return errors.New("No content")
	}

	return nil
    
}

func extractMatrix(w http.ResponseWriter, r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	err = validatePayload(records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// //


// StringifyMatrix returns a string representation of the ingested matrix. Used by the "/echo" endpoint
func StringifyMatrix(records [][]string) string {
    var s []string
    for i := range records {
        for ii, n := range records[i] {
            s = append(s, fmt.Sprintf("%s", n))			
			
			if ii == (len(records) - 1) {
				s = append(s, fmt.Sprintf("%s", "\n"))
			} else {
				s = append(s, fmt.Sprintf("%s", ","))
			}
        }
    }
    return strings.Join(s, "")
}





func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {		
		records, err := extractMatrix(w, r)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
			return	
		}
		
		response := StringifyMatrix(records)
		fmt.Fprint(w, response)
	})
	http.ListenAndServe(":8080", nil)
}
