package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed *.html
var FS embed.FS

type contactForm struct {
	Message string
	Error   string
}

func main() {
	tmpl := template.Must(template.ParseFS(FS, "form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, contactForm{})
	})

	http.HandleFunc("/send-message", func(w http.ResponseWriter, r *http.Request) {
		if len(r.FormValue("message")) < 10 {
			form := contactForm{
				Message: r.FormValue("message"),
				Error:   "Message too short",
			}

			tmpl.ExecuteTemplate(w, "contactform", form)
			return
		}

		http.Redirect(w, r, "/success", http.StatusFound)
	})

	http.HandleFunc("/success", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Message sent!"))
	})

	http.ListenAndServe(":8080", nil)
}
