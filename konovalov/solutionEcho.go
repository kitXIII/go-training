package konovalov

import (
	"io/ioutil"
	"log"
	"net/http"
)

// RunEchoServer starts the echo server
func RunEchoServer() {
	http.HandleFunc("/", myEchoHandler)
	log.Print(http.ListenAndServe(":8080", nil))
}

// myEchoHandler handle http request
func myEchoHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var body []byte

	isNotPostMethodError := httpError{message: "Expected POST method", status: 404}
	serverError := httpError{message: "Internal server error", status: 500}

	if r.Method != "POST" {
		responseWithError(w, isNotPostMethodError)
		return
	}

	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		responseWithError(w, serverError)
		log.Print(err)
		return
	}

	_, err = w.Write(body)

	if err != nil {
		responseWithError(w, serverError)
		log.Print(err)
		return
	}
}

type httpError struct {
	message string
	status  int
}

func responseWithError(w http.ResponseWriter, err httpError) {
	http.Error(w, err.message, err.status)
}
