package main

import (
	"fmt"
	"log"
	"net/http"
)

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
}

func main() {
	var s Struct
	err := http.ListenAndServe("localhost:4000", s)
	if err != nil {
		log.Fatal(http.ListenAndServe("localhost:4000", nil))
	}
}
