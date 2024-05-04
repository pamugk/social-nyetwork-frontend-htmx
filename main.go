package main

import (
    "html/template"
    "net/http"
)

func main() {
    templates := template.Must(template.ParseFiles("templates/register.html"))
    http.HandleFunc("/register", func (w http.ResponseWriter, r *http.Request) {
        templates.ExecuteTemplate(w, "register.html", nil)
    })

    http.ListenAndServe(":8080", nil)
} 
