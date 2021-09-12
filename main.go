package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			log.Panicln(err.Error())
		}
	})

	log.Fatalln(http.ListenAndServe(":1234", nil))
}
