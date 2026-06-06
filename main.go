package main

import (
	"ascii-web/handler"
	"log"
	"net/http"
)

func main() {

	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler.Home)

	log.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Faild to start server", err)
	}
}
