package http

import (
	"fmt"
	"net/http"
	"sync"
)

var message = ""
var mu sync.Mutex
var syncChannel = make(chan struct{})

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	message := getMessage()
	fmt.Fprintln(w, message)
}

func getMessage() string {
	return message
}

func addHello() {
	mu.Lock()
	defer mu.Unlock()
	message += "Hello"
	syncChannel <- struct{}{}
}

func addWorld() {
	<-syncChannel
	mu.Lock()
	defer mu.Unlock()
	message += " World!"
}

func StartServer() {
	http.HandleFunc("/hello", HelloWorldHandler)
	go addHello()
	go addWorld()
	http.ListenAndServe(":8080", nil)
}
