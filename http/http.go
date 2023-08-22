// http.go
package http

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request, message *string) {
	fmt.Fprintln(w, *message)
}

func StartServer(message *string) {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		HelloWorldHandler(w, r, message)
	})
	http.ListenAndServe(":8080", nil)
}
