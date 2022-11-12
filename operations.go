package main

import (
	"fmt"
	"strconv"
	"strings"
)

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


// SumMatrix returns the total sum of the elements in the matrix. Used by the "/sum" endpoint
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


// MultiplyMatrix returns the product of the integers in the matrix. Used by the "/multiply" endpoint
func MultiplyMatrix(records [][]string) (int, error) {
	var problem error
    product := 1
    for i := range records {
        for _, n := range records[i] {
			num, err := strconv.Atoi(n)
			if err != nil {
				problem = err
				break
			}
            product = product * num
        }
    }

	if problem != nil {
		return 0, problem
	}

    return product, nil
}
