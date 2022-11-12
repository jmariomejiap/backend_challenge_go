package main

import (
	"encoding/csv"
	"errors"
	"net/http"
	"strings"
)

//////////////////
// Shared Utils //
//////////////////

// validatePayload determines if the payload (matrix) is valid
func validatePayload(payload [][]string) error {
	if len(payload) == 0 {
		return errors.New("No content")
	}
	if len(payload) != len(payload[0]) {
		return errors.New("Matrix is not square. Rows and columns must have equal length")
	}
	// verify all elements in the matrix are integers
	_, err := SumMatrix(payload)
	if err != nil {
		if strings.Contains(err.Error(), "strconv.Atoi") {
			return errors.New("Matrix can only contain integers")
		}
		return err
	}

	return nil
}

// extractMatrix is responsible for reading a csv file and returning the Matrix
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
