package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed *.html
var FS embed.FS

func main() {
	var (
		tmpl    = template.Must(template.ParseFS(FS, "counter.html"))
		counter = 0
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, counter)
	})

	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		counter++
		tmpl.ExecuteTemplate(w, "counter", counter)
	})

	http.HandleFunc("/down", func(w http.ResponseWriter, r *http.Request) {
		counter--
		tmpl.ExecuteTemplate(w, "counter", counter)
	})

	http.ListenAndServe(":8080", nil)
}
