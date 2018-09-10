package handler

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	t := template.Must(template.ParseFiles("template/index.html"))

	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Fatal(err)
	}
}
