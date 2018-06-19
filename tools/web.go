package tools

import (
	"fmt"
	"net/http"
)

func ExecWeb() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/image", webimage)
	http.ListenAndServe("localhost:8000", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! url=%s", r.URL)
}

func webimage(w http.ResponseWriter, r *http.Request) {
	Lissajous(w)
}
