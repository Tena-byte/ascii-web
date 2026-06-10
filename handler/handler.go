package handler

import (
	"ascii-web/ascii"
	"html/template"
	"net/http"
	"strings"
)

type Ascii struct {
	Result string
}

func Home(w http.ResponseWriter, r *http.Request) {

	var tmpl = template.Must(template.ParseFiles("template/index.html"))

	data := Ascii{}

	if r.Method == http.MethodPost {

		text := r.FormValue("text")

		text = strings.NewReplacer("\r\n", "\n").Replace(text)
		text = strings.ReplaceAll(text, "\n", " ")

		
		banner := r.FormValue("banner")

		font := ascii.LoadBanner(banner)
		data.Result = ascii.Render(text, font)

	}

	tmpl.Execute(w, data)

}
