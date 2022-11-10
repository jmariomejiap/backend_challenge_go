package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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
	if len(paylaod) != len(paylaod[0]) {
		return errors.New("Matrix is not square. Rows and columns must be equal")
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

// InvertMatrix returns a string representation of the ingested matrix where the order is changed from row to columns. Used by the "/invert" endpoint
func InvertMatrix(records [][]string) string {
    var inverted []string
    for i := range records {
        for ii := range records[i] {
            inverted = append(inverted, fmt.Sprintf("%s", records[ii][i]))			
			
			if ii == (len(records) - 1) {
				inverted = append(inverted, fmt.Sprintf("%s", "\n"))
			} else {
				inverted = append(inverted, fmt.Sprintf("%s", ","))
			}
        }
    }
    return strings.Join(inverted, "")
}

// FlatMatrix returns the ingested matrix as a 1 line string, with values separated by commas. Used by the "/flatten" endpoint
func FlatMatrix(records [][]string) string {
	matrixSize := len(records)
    var s []string
    for i := range records {
        for ii, n := range records[i] {
            s = append(s, fmt.Sprintf("%s", n))
			s = append(s, fmt.Sprintf("%s", ","))
			if i != (matrixSize - 1) && ii != (matrixSize - 1) {
			}
        }
    }

	result := strings.Join(s, "")
    return result[:len(result) - 1] // trim last ","
}

// SumMatrix returns the total sum of the elements in the matrix
func SumMatrix(records [][]string) (int, error) {
	var problem error
    total := 0
    for i := range records {
        for _, n := range records[i] {
			num, err := strconv.Atoi(n)
			if err != nil {
				problem = err
				break
			}
            total += num
        }
    }

	if problem != nil {
		return 0, problem
	}

    return total, nil
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
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {		
		records, err := extractMatrix(w, r)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
			return	
		}
		
		response := InvertMatrix(records)
		fmt.Fprint(w, response)
	})
	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {		
		records, err := extractMatrix(w, r)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("Oops, something went wrong:\n%s", err.Error())))
			return	
		}
		
		response := FlatMatrix(records)
		fmt.Fprint(w, response, "\n")
	})
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {		
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
	})
	http.ListenAndServe(":8080", nil)
}
