package main

import (
	"ascii-web/handler"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler.Home)

	log.Println("server is running on http://localhost:5000")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal("server not running", err)
	}
}
