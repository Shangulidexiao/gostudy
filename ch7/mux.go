package ch7

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type yuan float64

type good map[string]yuan

func (y yuan) String() string {
	return fmt.Sprintf("%.2f", y)
}

func muxServer() {
	goods := good{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(goods.list))
	mux.Handle("/price", http.HandlerFunc(goods.price))
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func handleF() {

	fmt.Println(errors.New("EOF") == errors.New("EOF"))
	os.Exit(1)
	goods := good{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", http.HandlerFunc(goods.list))
	http.HandleFunc("/price", http.HandlerFunc(goods.price))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func (g good) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range g {
		fmt.Fprintf(w, "%s :%s\n", item, price)
	}
}

func (g good) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := g[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such good:%s\n", item)
		return
	}

	fmt.Fprintf(w, "%s :%s\n", item, db[item])
}
