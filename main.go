package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	templates := template.Must(template.ParseFiles("templates/register.html"))
	http.HandleFunc("/register", func (w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "register.html", nil)
	} else if r.Method == http.MethodPost {
		_, err := http.Post(os.Getenv("SERVICE_URL"))
		if err != nil {
			log.Fatalln(err)
		}
	}
	})

	http.ListenAndServe(":8081", nil)
} 
