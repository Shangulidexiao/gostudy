package ch7

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float64

type database map[string]dollars

var db = database{"shoes": 50, "socks": 5}

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func shop() {
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route := req.URL.Path
	switch route {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s :%s\n", item, price)
		}
	case "/price":
		good := req.URL.Query().Get("item")
		if _, ok := db[good]; !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such good:%s\n", good)
			return
		}

		fmt.Fprintf(w, "%s :%s\n", good, db[good])
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page:%s\n", req.URL)
	}

}
