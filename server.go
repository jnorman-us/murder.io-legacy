package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./server/static"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
