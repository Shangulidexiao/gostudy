// Server is a minimal "echo" server
package tools

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu1 sync.Mutex
var count int

func EchoUrl() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu1.Lock()
	count++
	mu1.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.RawPath)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu1.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu1.Unlock()
}
