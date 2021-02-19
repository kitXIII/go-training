package konovalov

import (
	"io/ioutil"
	"net/http"
)

// RunEchoServer starts the echo server
func RunEchoServer() {
	http.HandleFunc("/", myEchoHandler)
	http.ListenAndServe(":8080", nil)
}

// myEchoHandler handle http request
func myEchoHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte

	isNotPostMethodError := CustomError{message: "Expected POST method", status: 404}
	serverError := CustomError{message: "Internal server error", status: 500}

	if r.Method != "POST" {
		responseError(w, isNotPostMethodError)
		return
	}

	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		responseError(w, serverError)
		return
	}

	_, err = w.Write(body)

	if err != nil {
		responseError(w, serverError)
		return
	}
}

// CustomError implements error interface, uses custom message
type CustomError struct {
	message string
	status  int
}

func (e CustomError) Error() string {
	return e.message
}

// Common error handling
func responseError(w http.ResponseWriter, err CustomError) {
	http.Error(w, err.Error(), err.status)
}
