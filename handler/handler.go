package handler

import (
	"ascii-web/ascii"
	"html/template"
	"net/http"
)

type Ascii struct {
	Result string
}

func Home(w http.ResponseWriter, r *http.Request) {

	var tmpl = template.Must(template.ParseFiles("./template/index.html"))

	data := Ascii{}

	if r.Method == http.MethodPost {

		textarea := r.FormValue("text")
		banner := r.FormValue("banner")

		font := ascii.LaodBanner(banner)
		data.Result = ascii.Render(textarea, font)

	}

	tmpl.Execute(w, data)

}
