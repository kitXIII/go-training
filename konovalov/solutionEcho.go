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
	var isNotPostMethodError NotPostMethodErrorType

	if r.Method != "POST" {
		handleError(w, isNotPostMethodError)
		return
	}

	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		handleError(w, err)
		return
	}

	_, err = w.Write(body)

	if err != nil {
		handleError(w, err)
		return
	}
}

// Common error handling
func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), 500)
}

// NotPostMethodErrorType implements error interface, but uses custom message
type NotPostMethodErrorType struct{}

func (e NotPostMethodErrorType) Error() string {
	return "Expected POST method"
}
