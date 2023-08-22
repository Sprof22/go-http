package http

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var message = ""
var mu sync.Mutex

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	message := getMessage() //get the message
	fmt.Fprintln(w, message)
}

func getMessage() string {
	return message
}

func addHello() {
	time.Sleep(5 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	message += "Hello"
}

func addWorld() {
	time.Sleep(10 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	message += " world!"
}

func StartServer() {
	http.HandleFunc("/hello", HelloWorldHandler)
	go addHello()
	go addWorld()
	http.ListenAndServe(":8080", nil)
}
