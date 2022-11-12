package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

/////////////////
// test helper //
/////////////////

// buildPayloadAndWriter is a reusable function that reads a csv file and return the necessary elements to build the request
func buildPayloadAndWriter(matrix string) (*bytes.Buffer, *multipart.Writer, error) {
	payload := &bytes.Buffer{}	
	writer := multipart.NewWriter(payload)
	
	pwd, _ := os.Getwd()
	file, errFile1 := os.Open(pwd + "/" + matrix)
	if errFile1 != nil {
		fmt.Println(errFile1.Error())
		return nil, nil, errFile1
	}
	defer file.Close()

	part1, errFile1 := writer.CreateFormFile("file",filepath.Base( pwd + matrix))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return nil, nil, errFile1
	}
	err := writer.Close()		
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return payload, writer, nil
}


func TestHomeRoute(t *testing.T) {
	t.Run("should return the servers description", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		HomeHandlerFunc(response, request)

		got := response.Body.String()
		want := Description

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}


func TestEchoRoute(t *testing.T) {
	t.Run("should fail if no payload is sent", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/echo", nil)		
		response := httptest.NewRecorder()

		EchoHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nrequest Content-Type isn't multipart/form-data"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix is not square", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-not-square.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/echo", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		EchoHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix is not square. Rows and columns must have equal length"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix contains something else besides integer", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-letters.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/echo", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		EchoHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix can only contain integers"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should successfully ECHO the matrix", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/echo", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		EchoHandlerFunc(response, request)

		got := response.Body.String()
		want := "1,2,3\n4,5,6\n7,8,9\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}


func TestInvertRoute(t *testing.T) {
	t.Run("should fail if no payload is sent", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/invert", nil)		
		response := httptest.NewRecorder()

		InvertMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nrequest Content-Type isn't multipart/form-data"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix is not square", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-not-square.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/invert", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		InvertMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix is not square. Rows and columns must have equal length"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix contains something else besides integer", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-letters.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/invert", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		InvertMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix can only contain integers"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should successfully INVERT the matrix", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/invert", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		InvertMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "1,4,7\n2,5,8\n3,6,9\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}


func TestFlattenRoute(t *testing.T) {
	t.Run("should fail if no payload is sent", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/flatten", nil)		
		response := httptest.NewRecorder()

		FlattenMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nrequest Content-Type isn't multipart/form-data"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix is not square", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-not-square.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/flatten", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		FlattenMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix is not square. Rows and columns must have equal length"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix contains something else besides integer", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-letters.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/flatten", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		FlattenMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix can only contain integers"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should successfully FLATTEN the matrix", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/flatten", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		FlattenMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "1,2,3,4,5,6,7,8,9\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}


func TestSumRoute(t *testing.T) {
	t.Run("should fail if no payload is sent", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/sum", nil)		
		response := httptest.NewRecorder()

		SumMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nrequest Content-Type isn't multipart/form-data"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix is not square", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-not-square.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/sum", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		SumMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix is not square. Rows and columns must have equal length"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix contains something else besides integer", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-letters.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/sum", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		SumMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix can only contain integers"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should successfully SUM the matrix", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/sum", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		SumMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "45\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

	
func TestMultiplyRoute(t *testing.T) {
	t.Run("should fail if no payload is sent", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/multiply", nil)		
		response := httptest.NewRecorder()

		MultiplyMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nrequest Content-Type isn't multipart/form-data"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix is not square", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-not-square.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/multiply", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		MultiplyMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix is not square. Rows and columns must have equal length"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should fail if matrix contains something else besides integer", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix-letters.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/multiply", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		MultiplyMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "Oops, something went wrong:\nMatrix can only contain integers"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should successfully MULTIPLY the matrix", func(t *testing.T) {
		payload, writer, err := buildPayloadAndWriter("matrix.csv")
		if err != nil {
			t.Errorf("Something went wrong running the test %s", err.Error())
			return
		}
		
		request, _ := http.NewRequest(http.MethodGet, "/multiply", payload)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		response := httptest.NewRecorder()

		MultiplyMatrixHandlerFunc(response, request)

		got := response.Body.String()
		want := "362880\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}