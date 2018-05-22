package tools

import (
	"fmt"
	"log"
	"net/http"
)

func PraseForm() {
	http.HandleFunc("/", echoImage)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func Header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]: %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for n, v := range r.Form {
		fmt.Fprintf(w, "%s: %s\n", n, v)
	}
}

func echoImage(w http.ResponseWriter, r *http.Request) {
	Lissajous(w)
}
