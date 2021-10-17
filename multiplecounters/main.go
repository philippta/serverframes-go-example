package main

import (
	"embed"
	"net/http"
	"text/template"
)

//go:embed *.html
var FS embed.FS

type Counters struct {
	Counter1 int
	Counter2 int
}

func main() {
	tmpl := template.Must(template.ParseFS(FS, "multiplecounters.html"))

	counters := Counters{Counter1: 0, Counter2: 0}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, counters)
	})

	http.HandleFunc("/increment1", func(w http.ResponseWriter, r *http.Request) {
		counters.Counter1++
		tmpl.ExecuteTemplate(w, "counter1", counters.Counter1)
	})

	http.HandleFunc("/increment2", func(w http.ResponseWriter, r *http.Request) {
		counters.Counter2++
		tmpl.ExecuteTemplate(w, "counter2", counters.Counter2)
	})

	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {
		counters.Counter1++
		counters.Counter2++
		tmpl.ExecuteTemplate(w, "counter1", counters.Counter1)
		tmpl.ExecuteTemplate(w, "counter2", counters.Counter2)
	})

	http.ListenAndServe(":8080", nil)
}
