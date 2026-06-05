package handler

import (
	"net/http"
	"ascii-web/ascii"
)
func Generate(w http.ResponseWriter, r *http.Request) {

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	result := ascii.Generate(text, banner)

	data := struct {
		Result string
	}{
		Result: result,
	}

	tmpl.Execute(w, data)
}
