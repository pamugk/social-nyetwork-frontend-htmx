package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/register", func (w http.ResponseWriter, r *http.Request) {
    })

    http.ListenAndServe(":80", nil)
} 
