package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string
	}{
		Name: "Ritesh Patel",
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
		return
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handlerHome)
	http.ListenAndServe(":3000", nil)
}
