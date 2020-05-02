package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

var count int

func main() {
	http.HandleFunc("/inc", increment)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Panic(err)
	}
}

func increment(w http.ResponseWriter, r *http.Request) {
	count++
	if _, err := io.WriteString(w, strconv.Itoa(count)); err != nil {
		log.Panic(err)
	}
}
