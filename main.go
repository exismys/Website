package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Handle routes
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/blog/", handlerBlog)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listening at http://localhost:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	blogTitles := getBlogList()

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
		return
	}

	t.Execute(w, blogTitles)
}

func handlerBlog(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Printf("Blog handler called with path %s\n", path)
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		http.NotFound(w, r)
		return
	}

	blogFile := fmt.Sprintf("blogs/%s.html", parts[2])
	blog := getBlogPost(blogFile)

	t, err := template.ParseFiles("templates/blog.html")
	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Fprintf(w, "Unable to load template")
		return
	}

	t.Execute(w, blog)
}
