package konovalov

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// RunEchoServer starts the echo server
func RunEchoServer() {
	http.HandleFunc("/", myEchoHandler)
	http.ListenAndServe(":8080", nil)
}

func myEchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, "Expected POST method")
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(body)
}
