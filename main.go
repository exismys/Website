package main

import (
	"fmt"
	"html/template"
	"net/http"
  "os"
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
  port := os.Getenv("PORT")
  if port == "" {
      port = "8080" // default port if not set
  }
  fmt.Printf("Listening to port %s", port)
  http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
