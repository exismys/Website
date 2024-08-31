package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	blogTitles, err := getBlogList()
	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
		return
	}

	t.Execute(w, blogTitles)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlerHome)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listening to port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
