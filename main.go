package main

import (
	"go-http/http"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var message = ""
var mu sync.Mutex

func main() {
	wg.Add(2)
	go addHello(&message)
	go addWorld(&message)

	wg.Wait()

	http.StartServer(&message)
}

func addHello(msg *string) {
	time.Sleep(5 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	*msg += "Hello"
	wg.Done()
}

func addWorld(msg *string) {
	time.Sleep(10 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	*msg += " world!"
	wg.Done()
}
