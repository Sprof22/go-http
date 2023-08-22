package http

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func StartServer() {
	http.HandleFunc("/hello", HelloWorldHandler)
	http.ListenAndServe(":8080", nil)
}
